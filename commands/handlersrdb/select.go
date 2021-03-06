package handlersrdb

import (
	"database/sql"

	"github.com/arteev/dsql/db"
	"github.com/arteev/dsql/format"
	"github.com/arteev/dsql/rdb/action"
	"github.com/arteev/dsql/rdb/sqlcommand"
	"github.com/arteev/logger"
	"github.com/nsf/termbox-go"

	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/arteev/dsql/dataset"
	"github.com/arteev/dsql/parameters/parametergetter"
	"github.com/arteev/dsql/rowgetter"
	"github.com/arteev/dsql/tabex"
	"github.com/arteev/fmttab"
	"github.com/arteev/fmttab/columns"
)

type headerTable struct {
	Code    string //Code of DB
	Columns []string
	Rows    int
}
type dataTable struct {
	Code string
	Data map[string]interface{}
}

type chanHeader chan *headerTable
type chanData chan *dataTable

func foundInSlice(strs columns.Columns, value string) bool {
	for _, s := range strs {
		if s.Name == value {
			return true
		}
	}
	return false
}

//formatRaw returns formatted string for output as Raw
func formatRaw(ctx *action.Context, f *format.Format, row, line int, columns []string, data map[string]interface{}) (string, error) {
	mask := f.RawString()
	if mask == "" {
		mask = "{COLUMN}:\"{VALUE}\";"
	}
	aliasSeparator := ctx.GetDef("sepalias", ": ").(string)
	alias := data["_CODE_"].(string)
	buf := &bytes.Buffer{}
	for c, col := range columns {
		val, ok := data[col]
		if ok {
			currentOut := strings.Replace(mask, "{COLUMN}", col, -1)
			currentOut = strings.Replace(currentOut, "{VALUE}", fmt.Sprintf("%v", val), -1)
			currentOut = strings.Replace(currentOut, "{COLNUM}", strconv.Itoa(c+1), -1)
			currentOut = strings.Replace(currentOut, "{LINE}", strconv.Itoa(line), -1)
			currentOut = strings.Replace(currentOut, "{ROW}", strconv.Itoa(row), -1)
			if _, err := buf.WriteString(currentOut); err != nil {
				return "", err
			}
		}
	}
	return alias + aliasSeparator + buf.String(), nil
}

//SelectBefore trigger before for select action
func SelectBefore(dbs []db.Database, ctx *action.Context) error {
	// Prepare data in ctx.datasets
	logger.Trace.Println("SelectBefore")
	immediate := ctx.GetDef("immediate", false).(bool)
	fm := ctx.Get("format").(*format.Format)
	//logger.Trace.Printf("%#v\n", fm)

	datasets := dataset.NewColllection()
	chanHdr := make(chanHeader)
	chandata := make(chanData)
	chanDone := make(chan bool)

	ctx.Set("chanheader", chanHdr)
	ctx.Set("chandata", chandata)
	ctx.Set("chandone", chanDone)
	ctx.Set("datasets", datasets)

	line := 0
	go func() {
		for {
			select {
			case hdr := <-chanHdr:
				ds := datasets.GetOrCreateDataset(hdr.Code)
				ds.AddColumns(hdr.Columns...)
			case cudata := <-chandata:
				line++
				ds := datasets.GetOrCreateDataset(cudata.Code)
				ds.Append(cudata.Data)
				if fm.Name() == "raw" && immediate {
					data, err := formatRaw(ctx, fm, ds.RowsCount(), line, ds.GetColumnsNames(), cudata.Data)
					if err != nil {
						logger.Error.Println(err)
					} else {
						fmt.Println(data)
					}
				}
			case <-chanDone:
				logger.Trace.Println("SelectBefore do done")
				return
			}
		}
	}()

	return nil
}

func doOutputTable(dbs []db.Database, ctx *action.Context) error {
	datasets := ctx.Get("datasets").(*dataset.CollectionDataset)
	table := fmttab.New("", fmttab.BorderThin, nil)
	table.AddColumn("_CODE_", 10, fmttab.AlignLeft)
	ctx.Set("table", table)

	for _, col := range datasets.GetUniqueColumnsNames() {
		table.AddColumn(col, 15, fmttab.AlignLeft)
	}
	for _, ds := range datasets.GetDatasets() {
		for _, row := range ds.Rows {
			table.AppendData(row.GetDataMap())
		}
	}

	pget := ctx.Get("params").(parametergetter.ParameterGetter)

	if pget.GetDef(parametergetter.AutoFitWidthColumns, true).(bool) {
		//todo: move into fmttab
		cols := table.Columns.ColumnsVisible()
		for c, col := range cols {
			max := utf8.RuneCountInString(col.Name)
			for i := 0; i < len(table.Data); i++ {
				val, ok := table.Data[i][col.Name]

				if ok && val != nil {
					fval := fmt.Sprintf("%v", val)
					l := utf8.RuneCountInString(fval)
					if l > max {
						max = l
					}
				}
			}
			if max != 0 {
				cols[c].Width = max
			}
		}
	}

	switch pget.GetDef(parametergetter.BorderTable, "").(string) {
	case "Thin":
		table.SetBorder(fmttab.BorderThin)
	case "Double":
		table.SetBorder(fmttab.BorderDouble)
	case "None":
		table.SetBorder(fmttab.BorderNone)
	case "Simple":
		table.SetBorder(fmttab.BorderSimple)
	}

	if err := tabex.SetTableSubFormat(ctx, table); err != nil {
		return err
	}

	if pget.GetDef(parametergetter.Fit, true).(bool) {
		if e := termbox.Init(); e != nil {
			return e
		}
		tw, _ := termbox.Size()
		table.AutoSize(true, tw)
		termbox.Close()
	}

	if _, err := table.WriteTo(os.Stdout); err != nil {
		return err
	}
	return nil
}

func getStrByIdx(params []string, idx int) (string, bool) {
	if len(params) <= idx {
		return "", false
	}
	return params[idx], true
}

func fillDatasetsByErrors(datasets *dataset.CollectionDataset, dbs []db.Database, ctx *action.Context) error {
	for _, d := range dbs {
		localCtx := ctx.Get("context" + d.Code).(*action.Context)
		if !localCtx.Get("success").(bool) {
			ds := datasets.GetOrCreateDataset(d.Code)
			ds.Error = true
			err := localCtx.Snap.Error()
			if err != nil {
				ds.TextError = err.Error()
			} else {
				ds.TextError = "WARN:Unknown error"
				logger.Warn.Println("WARN: Unknown error (fillDatasetsByErrors)")

			}
		}
	}
	return nil
}

func doOutputJSON(dbs []db.Database, ctx *action.Context) error {
	datasets := ctx.Get("datasets").(*dataset.CollectionDataset)
	indent := ctx.GetDef("indent", "\t").(string)
	if err := fillDatasetsByErrors(datasets, dbs, ctx); err != nil {
		return err
	}

	fm := ctx.Get("format").(*format.Format)
	filename, ok := fm.Root().Get("file")
	if !ok {
		_, err := datasets.WriteJSON(os.Stdout, indent)
		return err
	}
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = datasets.WriteJSON(f, indent)
	return err
}

func doOutputRaw(dbs []db.Database, ctx *action.Context) error {
	datasets := ctx.Get("datasets").(*dataset.CollectionDataset)

	fm := ctx.Get("format").(*format.Format)
	line := 0
	for _, ds := range datasets.GetDatasets() {
		for _, row := range ds.Rows {
			line++
			data, err := formatRaw(ctx, fm, ds.RowsCount(), line, ds.GetColumnsNames(), row.GetDataMap())
			if err != nil {
				return err
			}
			fmt.Println(data)
		}
	}

	return nil
}

func doOutputXML(dbs []db.Database, ctx *action.Context) error {
	datasets := ctx.Get("datasets").(*dataset.CollectionDataset)
	indent := ctx.GetDef("indent", "    ").(string)
	if err := fillDatasetsByErrors(datasets, dbs, ctx); err != nil {
		return err
	}
	fm := ctx.Get("format").(*format.Format)
	filename, ok := fm.Root().Get("file")
	if !ok {
		datasets.WriteXML(os.Stdout, indent)
		return nil
	}
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = datasets.WriteXML(f, indent)
	return err
}

//SelectAfter trigger after for select action
func SelectAfter(dbs []db.Database, ctx *action.Context) error {
	done := ctx.Get("chandone")
	immediate := ctx.GetDef("immediate", false).(bool)
	logger.Trace.Println("SelectAfter")
	if done != nil {
		done.(chan bool) <- true
	}
	fm := ctx.Get("format").(*format.Format)
	switch fm.Name() {
	case "table":
		return doOutputTable(dbs, ctx)
	case "json":
		return doOutputJSON(dbs, ctx)
	case "xml":
		return doOutputXML(dbs, ctx)
	case "raw":
		if !immediate {
			return doOutputRaw(dbs, ctx)
		}
	}
	return nil
}

//SelectError trigger error for select action
func SelectError(dbs []db.Database, ctx *action.Context) error {

	logger.Trace.Println("Failed execute")
	if !ctx.GetDef("silent", false).(bool) {
		fmt.Println("All requests will fail.")
	}
	fm := ctx.Get("format").(*format.Format)
	switch fm.Name() {
	case "json", "xml":
		return SelectAfter(dbs, ctx)
	}

	done := ctx.Get("chandone")
	if done != nil {
		done.(chan bool) <- true
	}
	return nil
}

//Select - it action for select command
func Select(dbs1 db.Database, dsrc *sql.DB, cmd *sqlcommand.SQLCommand, ctx *action.Context) error {
	logger.Trace.Println("run select", dbs1.Code)

	timeout := ctx.GetDef("timeout", 0).(int)
	logger.Debug.Printf("run select timeout %d sec", timeout)

	var pint []interface{}

	localCtx := ctx.Get("context" + dbs1.Code).(*action.Context)
	if localCtx == nil {
		return fmt.Errorf("Could not get local context %s", dbs1.Code)
	}

	chanCancel := localCtx.Get("iscancel").(chan bool)
	for _, p := range cmd.Params {
		pint = append(pint, p)
	}

	/*tx,err:=dsrc.Begin()
	if err!=nil {
		return err
	}
	defer tx.Rollback()
	rw, err := tx.Query(cmd.Script, pint...)*/
	tx, err := dsrc.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	stmt, err := tx.Prepare(cmd.Script)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rw, err := tx.Query(cmd.Script, pint...)
	if err != nil {
		return err
	}

	defer func() {
		if err := rw.Close(); err != nil {
			panic(err)
		}
	}()

	cols, _ := rw.Columns()

	chanHdr := ctx.Get("chanheader")
	chandata := ctx.Get("chandata")
	if chanHdr != nil {
		chanHdr.(chanHeader) <- &headerTable{
			Code:    dbs1.Code,
			Columns: cols,
		}
	}

	rg := rowgetter.MustRowGetter(rw)
	for {
		select {
		case <-chanCancel:
			logger.Info.Println("run select canceled", dbs1.Code)
			return fmt.Errorf("cancel")
		default:
		}
		row, ok := rg.Next()
		if !ok {
			break
		}

		localCtx.IncInt("rowcount", 1)

		if /*format=="table" &&*/ chandata != nil {
			data := make(map[string]interface{})
			for i, r := range row {
				data[cols[i]] = r
			}
			data["_CODE_"] = dbs1.Code

			chandata.(chanData) <- &dataTable{
				Code: dbs1.Code,
				Data: data,
			}
		}

	}

	return nil
}
