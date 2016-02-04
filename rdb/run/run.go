package run

import (
	"github.com/arteev/dsql/db"

	"github.com/arteev/dsql/rdb/action"
	"github.com/arteev/dsql/rdb/sqlcommand"

	"sync"

	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/arteev/dsql/parameters"
	"github.com/arteev/dsql/parameters/parametergetter"
	"github.com/arteev/dsql/parameters/paramsreplace"
	"github.com/arteev/dsql/rdb"
	"github.com/arteev/logger"
)

var wg sync.WaitGroup

//Run concrete action for all databases
func Run(dbs []db.Database, sql *sqlcommand.SQLCommand, act action.Actioner, ctx *action.Context, pget parametergetter.ParameterGetter) (*action.Context, error) {
	logger.Trace.Println("rdb run")
	defer logger.Trace.Println("rdb run done")
	ctx.Set("params", pget)
	ctx.Set("silent", pget.GetDef(parametergetter.Silent, false).(bool))
	ctx.Snap.Start()

	colParams, err := parameters.GetInstance().All()
	if err != nil {
		return nil, err
	}
	ctx.Set("Params", colParams.Get())

	for _, d := range dbs {
		err := runItem(d, sql, act, ctx, pget)
		if err != nil {
			logger.Error.Println(err)
			//todo: в общий список ошибок
		}
	}
	wg.Wait()
	ctx.Snap.Done(nil)
	return ctx, nil
}

func runItem(d db.Database, s *sqlcommand.SQLCommand, act action.Actioner, ctx *action.Context, pget parametergetter.ParameterGetter) error {
	logger.Trace.Println("runItem")
	defer logger.Trace.Println(d.Code, "runItem done")
	if s != nil {
		logger.Trace.Println(d.Code, s.Script)
	}
	wg.Add(1)
	ctx.IncInt("exec", 1)
	params := ctx.Get("Params").([]parameters.Parameter)
	go func() {

		localCtx := action.NewContext()
		ctx.Set("context"+d.Code, localCtx)
		localCtx.Snap.Start()
		localCtx.Set("success", false)

		connectionString, e := paramsreplace.Replace(d.ConnectionString, params)
		if e != nil {
			ctx.IncInt("failed", 1)
			logger.Error.Println(e)
			return
		}
		logger.Debug.Println(d.Code, "Connection string:", connectionString)

		connection, err := rdb.Open(d.Engine, connectionString)
		defer wg.Done()
		if err != nil {
			ctx.IncInt("failed", 1)
			logger.Error.Println(d.Code, err)
			return
		}
		defer func() {
			if err := connection.Close(); err != nil {
				panic(err)
			}
		}()

		err = act(d, connection, s, ctx)
		if err != nil {
			ctx.IncInt("failed", 1)
			localCtx.Snap.Done(err)
			logger.Error.Println(d.Code, err)
			if !ctx.GetDef("silent", false).(bool) {
				fmt.Fprintf(os.Stdout, "%s: %s\n", d.Code, strings.Replace(err.Error(), "\n", " ", -1))
			}
			return
		}

		localCtx.Set("success", true)
		ctx.IncInt("success", 1)
		localCtx.Snap.Done(nil)
		runtime.Gosched()
	}()
	return nil
}
