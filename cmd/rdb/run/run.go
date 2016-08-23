package run

import (
	"github.com/arteev/dsql/cmd/db"

	"github.com/arteev/dsql/cmd/rdb/action"
	"github.com/arteev/dsql/cmd/rdb/sqlcommand"

	"sync"

	"fmt"
	"os"
	"runtime"
	"strings"

	"time"

	"github.com/arteev/dsql/cmd/parameters"
	"github.com/arteev/dsql/cmd/parameters/parametergetter"
	"github.com/arteev/dsql/cmd/parameters/paramsreplace"
	"github.com/arteev/dsql/cmd/rdb"
	"github.com/arteev/logger"
	"golang.org/x/net/context"
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

func runItem(d db.Database, s *sqlcommand.SQLCommand, doaction action.Actioner, ctx *action.Context, pget parametergetter.ParameterGetter) error {
	logger.Trace.Println("runItem")
	defer logger.Trace.Println(d.Code, "runItem done")
	if s != nil {
		logger.Trace.Println(d.Code, s.Script)
	}
	wg.Add(1)
	ctx.IncInt("exec", 1)
	params := ctx.Get("Params").([]parameters.Parameter)

	go func() {

		timeout := ctx.GetDef("timeout", 0).(int)
		defer wg.Done()

		var (
			ctxExec context.Context
			cancel  context.CancelFunc
		)
		ch := make(chan bool)
		if timeout > 0 {
			ctxExec, cancel = context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
		} else {
			ctxExec, cancel = context.WithCancel(context.Background())
		}

		defer cancel()
		localCtx := action.NewContext()

		go func() {

			defer func() {
				ch <- true
				close(ch)
			}()

			ctx.Set("context"+d.Code, localCtx)
			ctx.Set("iscancel", ch)
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
			if err != nil {
				ctx.IncInt("failed", 1)
				logger.Error.Println(d.Code, err)
				return
			}
			defer func() {
				if err := connection.Close(); err != nil {
					panic(err)
				} else {
					logger.Trace.Printf("%s disconnected", d.Code)
				}
			}()

			err = doaction(d, connection, s, ctx)
			if err != nil {
				if err.Error() != "cancel" {
					ctx.IncInt("failed", 1)
					localCtx.Snap.Done(err)
					logger.Error.Println(d.Code, err)
					if !ctx.GetDef("silent", false).(bool) {
						fmt.Fprintf(os.Stdout, "%s: %s\n", d.Code, strings.Replace(err.Error(), "\n", " ", -1))
					}
				}

				return
			}

			localCtx.Set("success", true)
			ctx.IncInt("success", 1)
			localCtx.Snap.Done(nil)
			runtime.Gosched()
		}()

		select {
		case <-ch:
			logger.Trace.Println("operation done w/o timeout")
			return
		case <-ctxExec.Done():
			err := ctxExec.Err()
			logger.Trace.Printf("operation done: %s\n", err)

			ctx.IncInt("failed", 1)
			localCtx.Snap.Done(err)
			logger.Error.Println(d.Code, err)

			//	ch <- true

			return
		}

	}()
	return nil
}
