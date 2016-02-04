package handlersrdb

import (
	"database/sql"

	"github.com/arteev/dsql/db"
	"github.com/arteev/dsql/rdb/action"
	"github.com/arteev/dsql/rdb/sqlcommand"
	"github.com/arteev/logger"

	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/arteev/dsql/rowgetter"
	"github.com/arteev/fmttab"
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

func foundInSlice(strs []*fmttab.Column, value string) bool {
	for _, s := range strs {
		if s.Name == value {
			return true
		}
	}
	return false
}

//formatRaw returns formatted string for output as Raw
func formatRaw(mask string, row, line int, columns []string, data map[string]interface{}) string {
	if mask == "" {
		mask = "{COLUMN}:\"{VALUE}\";"
	}
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
				panic(err)
			}
		}
	}
	return data["$CODE$"].(string) + ": " + buf.String()
}

//SelectBefore trigger before for select action
func SelectBefore(dbs []db.Database, ctx *action.Context) error {
	logger.Trace.Println("SelectBefore")
	format := ctx.Get("format")
	subformat := ctx.GetDef("subformat", "").(string)

	headers := make(map[string]*headerTable)
	chanHdr := make(chanHeader)
	chandata := make(chanData)
	chanDone := make(chan bool)

	var tab *fmttab.Table
	if format == "table" {
		tab = fmttab.New("", fmttab.BorderThin, nil)
		tab.AddColumn("$CODE$", 10, fmttab.AlignLeft)
		ctx.Set("table", tab)
	}
	ctx.Set("chanheader", chanHdr)
	ctx.Set("chandata", chandata)
	ctx.Set("chandone", chanDone)

	line := 0
	go func() {
		for {
			select {
			case hdr := <-chanHdr:
				h, ok := headers[hdr.Code]
				if !ok {
					h = &headerTable{}
					headers[hdr.Code] = h
				}
				h.Columns = hdr.Columns
				if format == "table" {
					for _, col := range hdr.Columns {
						if !foundInSlice(tab.Columns, col) {
							tab.AddColumn(col, 15, fmttab.AlignLeft)
						}
					}
				}
			case data := <-chandata:
				line++
				h, ok := headers[data.Code]
				if !ok {
					h = &headerTable{}
					headers[data.Code] = h
				}
				h.Rows++
				if format == "table" {
					tab.AppendData(data.Data)
				} else if format == "raw" {
					fmt.Println(formatRaw(subformat, h.Rows, line, h.Columns, data.Data))
				}
			case <-chanDone:
				logger.Trace.Println("SelectBefore do done")
				return
			}
		}
	}()

	return nil
}

//SelectAfter trigger after for select action
func SelectAfter(dbs []db.Database, ctx *action.Context) error {
	done := ctx.Get("chandone")
	logger.Trace.Println("SelectAfter")
	if done != nil {
		done.(chan bool) <- true
	}
	tab := ctx.Get("table")
	if tab != nil {
		table := tab.(*fmttab.Table)

		//autofit
		for c, col := range table.Columns {
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
				table.Columns[c].Width = max
			}
		}
		if _, err := table.WriteTo(os.Stdout); err != nil {
			return err
		}
	}
	return nil
}

//SelectError trigger error for select action
func SelectError(dbs []db.Database, ctx *action.Context) error {
	done := ctx.Get("chandone")
	if done != nil {
		done.(chan bool) <- true
	}
	logger.Trace.Println("Failed execute")
	if !ctx.GetDef("silent", false).(bool) {
		fmt.Println("All requests will fail.")
	}
	return nil
}

//Select - it action for select command
func Select(dbs1 db.Database, dsrc *sql.DB, cmd *sqlcommand.SQLCommand, ctx *action.Context) error {
	logger.Trace.Println("run select", dbs1.Code)
	var pint []interface{}
	for _, p := range cmd.Params {
		pint = append(pint, p)
	}
	rw, err := dsrc.Query(cmd.Script, pint...)
	if err != nil {
		return err
	}
	defer func() {
		if err := rw.Close(); err != nil {
			panic(err)
		}
	}()

	cols, _ := rw.Columns()

	localCtx := ctx.Get("context" + dbs1.Code).(*action.Context)
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
			data["$CODE$"] = dbs1.Code

			chandata.(chanData) <- &dataTable{
				Code: dbs1.Code,
				Data: data,
			}
		}

	}
	if err := rw.Err(); err != nil {
		return err
	}
	return nil
}
