package commands

import (
	"bufio"
	"fmt"
	"io"
	"net/url"
	"os"
	"strconv"

	"strings"

	"github.com/arteev/dsql/commands/handlersrdb"
	"github.com/arteev/dsql/db"
	"github.com/arteev/dsql/parameters"
	"github.com/arteev/dsql/parameters/parametergetter"
	"github.com/arteev/dsql/parameters/paramsreplace"
	"github.com/arteev/dsql/rdb"
	"github.com/arteev/dsql/rdb/action"
	"github.com/arteev/dsql/rdb/run"
	"github.com/arteev/dsql/rdb/sqlcommand"
	"github.com/arteev/logger"
	"github.com/urfave/cli"
)

//getDatabases return get enabled database with filter by name from args application
func getDatabases(cflags *cliFlags, r db.RepositoryDB) ([]db.Database, error) {

	uriList := make([]db.Database, 0)
	for i, v := range cflags.DatabasesURI() {
		engine := "firebirdsql"
		uri := v
		u, err := url.Parse(v)
		if err == nil && u.Scheme != "" && rdb.CheckCodeEngine(u.Scheme) == nil {
			engine = u.Scheme
			if engine != "postgres" {
				u.Scheme = ""
				uri = strings.TrimLeft(u.String(), "/")
			}
		}

		fmt.Printf("engine: %q, uri %q, %q", engine, v, uri)
		d := db.Database{
			ID:               -(i + 1),
			Code:             "udb" + strconv.Itoa(i+1),
			Enabled:          true,
			ConnectionString: uri,
			Engine:           engine,
			Tags:             make([]*db.Tag, 0),
		}
		uriList = append(uriList, d)
	}

	if (len(uriList) != 0) && len(cflags.Databases()) == 0 {
		return uriList, nil
	}

	dbs, err := r.All()
	if err != nil {
		return nil, err
	}
	dbs.AddFilterEnabled()
	cflags.ApplyTo(dbs)
	for _, e := range cflags.Engines() {
		if err := rdb.CheckCodeEngine(e); err != nil {
			return nil, err
		}
	}
	result := dbs.Get()
	result = append(result, uriList...)
	return result, nil
}

func parseSQLFromArgs(ctx *cli.Context) *sqlcommand.SQLCommand {
	var sqlText string
	if !ctx.IsSet("sql") {
		//Trying from stdin
		fi, err := os.Stdin.Stat()
		if err != nil {
			panic(err)
		}
		if fi.Mode()&os.ModeNamedPipe == 0 {
			return nil
		}
		bio := bufio.NewReader(os.Stdin)
		sqlText, err = bio.ReadString(0)
		if err != nil && err != io.EOF {
			panic(err)
		}
	} else {
		sqlText = ctx.String("sql")
	}
	//Prepare parameters
	colParams, e := parameters.GetInstance().All()
	if e != nil {
		panic(e)
	}
	params := colParams.Get()
	paramsArgs := ctx.StringSlice("param")
	for i := 0; i < len(paramsArgs); i++ {

		newparam, e := paramsreplace.Replace(paramsArgs[i], params)
		if e != nil {
			logger.Error.Println(e)
		} else {
			paramsArgs[i] = newparam
		}
	}
	return sqlcommand.New(sqlText, paramsArgs)
}

func parseOthersFlagsForRunContext(ctx *cli.Context, ctxRun *action.Context) error {
	if ctx.IsSet("format") {
		format := ctx.String("format")
		subformat := ""
		//TODO: refactor it!
		if strings.Contains(format, "raw:") {
			subformat = format[len("raw:"):]
			format = "raw"
		}
		if strings.Contains(format, "table:") {
			subformat = format[len("table:"):]
			format = "table"
		}
		if strings.Contains(format, "json:") {
			subformat = format[len("json:"):]
			format = "json"
		}
		if strings.Contains(format, "xml:") {
			subformat = format[len("xml:"):]
			format = "xml"
		}

		switch format {
		case "table", "raw", "json", "xml":
			ctxRun.Set("format", format)
			ctxRun.Set("subformat", subformat)
			break
		default:
			return fmt.Errorf("Unknown format:%s", format)
		}
	} else {
		ctxRun.Set("format", "raw")
	}

	if ctx.IsSet("timeout") {
		ctxRun.Set("timeout", ctx.Int("timeout"))
	}
	if ctx.IsSet("commit") {
		ctxRun.Set("commit", ctx.Bool("commit"))
	}
	if ctx.IsSet("immediate") {
		ctxRun.Set("immediate", ctx.Bool("immediate"))
	}
	if ctx.IsSet("sepalias") {
		ctxRun.Set("sepalias", ctx.String("sepalias"))

	}
	if ctx.IsSet("indent") {
		ctxRun.Set("indent", ctx.String("indent"))
	}
	return nil
}

//
func createParametersGetter(ctx *cli.Context) parametergetter.ParameterGetter {
	return parametergetter.New(ctx, parameters.GetInstance())
}

type actionTriggerDBS func(dbs []db.Database, ctx *action.Context) error

func muxActionTriggers(trgs ...actionTriggerDBS) actionTriggerDBS {
	return func(dbs []db.Database, ctx *action.Context) error {
		for _, t := range trgs {
			err := t(dbs, ctx)
			if err != nil {
				return err
			}
		}
		return nil
	}
}

func doTrigger(a actionTriggerDBS, dbs []db.Database, ctx *action.Context) {
	if a == nil {
		return
	}
	if err := a(dbs, ctx); err != nil {
		panic(err)
	}
}

//commonActionDBS. Returns action for cli app bind with handler for current db item
func commonActionDBS(cflags *cliFlags, name string, a action.Actioner, sqlRequired bool, before, after, errtrg actionTriggerDBS) func(ctx *cli.Context) {
	return func(ctx *cli.Context) {
		logger.Trace.Println("command", name)
		defer logger.Trace.Println("command", name, "done")
		d := db.GetInstance()
		defer checkErr(d.Close)
		cflags.SetContext(ctx)

		paramGetter := createParametersGetter(ctx)

		var sc *sqlcommand.SQLCommand
		if sqlRequired {
			sc = parseSQLFromArgs(ctx)
			if sc == nil {
				panic(fmt.Errorf("SQL not found"))
			}
			logger.Debug.Printf("select sql:%q, %v\n", sc.Script, sc.Params)
		}

		contextRun := action.NewContext()

		if err := parseOthersFlagsForRunContext(ctx, contextRun); err != nil {
			panic(err)
		}

		dbsSource, err := getDatabases(cflags, d)
		if err != nil {
			panic(err)
		}
		if len(dbsSource) == 0 {
			panic(fmt.Errorf("databases not found"))
		}

		doTrigger(before, dbsSource, contextRun)
		if _, e := run.Run(dbsSource, sc, a, contextRun, paramGetter); e != nil {
			logger.Error.Println(e)
			doTrigger(errtrg, dbsSource, contextRun)
		} else {

			if success, ok := contextRun.Get("success").(int); !ok || success == 0 {
				doTrigger(errtrg, dbsSource, contextRun)
			} else {
				doTrigger(after, dbsSource, contextRun)
			}
		}

	}
}

//flagsForQuery - Define flags for query (select,exec)
func flagsForQuery(fs ...cli.Flag) []cli.Flag {
	result := []cli.Flag{
		cli.StringFlag{
			Name:  "sql",
			Usage: "sql text. use prefix @ from file. @/path/file.sql",
		},
		cli.StringSliceFlag{
			Name:  "param",
			Usage: "set params for query",
		},
		cli.StringFlag{
			Name:  "format",
			Usage: "Format output: raw[:subformat] (default) | table|xml|json.  Subformat - use templates: {ALIAS}{COLNUM}{ROW}{LINE}{COLUMN}{VALUE}  ",
			Value: "raw",
		},
		cli.BoolFlag{
			Name:  "immediate",
			Usage: "Whenever possible, output data directly (raw)",
		},
		cli.IntFlag{
			Name:  "timeout",
			Usage: "Timeout operation. Default 0",
			Value: 0,
		},
	}
	result = append(result, fs[:]...)
	return result
}

func combineFlags(flags ...cli.Flag) []cli.Flag {
	var result []cli.Flag
	result = append(result, flags[:]...)
	return result
}

//GetCommandsDBS returns the command for register in cli app
func GetCommandsDBS() []cli.Command {
	dbFilterFlags := newCliFlags(cliOption{
		Databases:        modeFlagMulti,
		ExcludeDatabases: modeFlagMulti,
		Engines:          modeFlagMulti,
		Tags:             modeFlagMulti,
		ExcludeTags:      modeFlagMulti,
	})
	flagsQuery := flagsForQuery(dbFilterFlags.Flags()...)
	return []cli.Command{
		cli.Command{
			Name:  "ping",
			Usage: "ping remote databases",
			Flags: flagsQuery,
			Action: commonActionDBS(dbFilterFlags, "ping", handlersrdb.PingHandler, false,
				nil,
				muxActionTriggers(handlersrdb.PrintStatistic),
				muxActionTriggers(handlersrdb.PrintStatistic)),
		},
		cli.Command{
			Name:  "select",
			Usage: "select from remote databases",
			Flags: append(flagsQuery,
				cli.BoolFlag{
					Name:  "fit",
					Usage: "use for fit table by width window of terminal",
				},
				cli.BoolFlag{
					Name:  "fitcolumns",
					Usage: "use for auto width columns by contents",
				},
				cli.StringFlag{
					Name:  "border",
					Usage: "set type of border table: Thin,Double,Simple or None. Default:Thin",
				},
				cli.StringFlag{
					Name:  "sepalias",
					Usage: "separator alias for output raw. default ': '",
				},
				cli.StringFlag{
					Name:  "indent",
					Usage: "indent output for format:xml,json. Default xml:4 spaces; json:\\t",
				}),
			Action: commonActionDBS(dbFilterFlags, "select", handlersrdb.Select, true,
				handlersrdb.SelectBefore,
				muxActionTriggers(handlersrdb.SelectAfter, handlersrdb.PrintStatisticQuery, handlersrdb.PrintStatistic, handlersrdb.WriteRetryFile),
				muxActionTriggers(handlersrdb.SelectError, handlersrdb.PrintStatisticQuery, handlersrdb.PrintStatistic, handlersrdb.WriteRetryFile)),
		},
		cli.Command{
			Name:  "exec",
			Usage: "execute sql command on the remote databases",
			Flags: append(flagsQuery,
				cli.BoolFlag{
					Name:  "commit",
					Usage: "Use flag for commit the transaction. Default: false(rollback)",
				},
			),
			Action: commonActionDBS(dbFilterFlags, "exec", handlersrdb.Exec, true,
				nil,
				muxActionTriggers(handlersrdb.PrintStatisticQuery, handlersrdb.PrintStatistic, handlersrdb.WriteRetryFile),
				muxActionTriggers(handlersrdb.PrintStatisticQuery, handlersrdb.PrintStatistic, handlersrdb.WriteRetryFile)),
		},
	}
}
