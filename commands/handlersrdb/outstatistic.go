package handlersrdb

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/arteev/dsql/db"
	"github.com/arteev/dsql/parameters/parametergetter"
	"github.com/arteev/dsql/rdb/action"
)

//PrintStatistic - print common statistic after execute action
func PrintStatistic(dbs []db.Database, ctx *action.Context) error {
	var buf bytes.Buffer
	pget := ctx.Get("params").(parametergetter.ParameterGetter)
	if pget.GetDef(parametergetter.Statistic, false).(bool) {
		mSec := ctx.Snap.Finished().Sub(ctx.Snap.Started()).Nanoseconds() / 1000 / 1000
		exec := ctx.GetDef("exec", 0).(int)
		success := ctx.GetDef("success", 0).(int)
		failed := ctx.GetDef("failed", 0).(int)
		buf.WriteString(fmt.Sprintf("Executed: %-4d Success:%-3d (%3.2f%%) Failed:%-3d \n", exec, success, float64(success)/float64(exec)*100, failed))
		buf.WriteString(fmt.Sprintf("Completed: %v msec", mSec))
		fmt.Println(buf.String())
	}
	return nil
}

//PrintStatisticQuery - print statistic for each database after execute action
func PrintStatisticQuery(dbs []db.Database, ctx *action.Context) error {
	pget := ctx.Get("params").(parametergetter.ParameterGetter)
	if pget.GetDef(parametergetter.QueryStatistic, false).(bool) {
		for _, d := range dbs {
			localCtx := ctx.Get("context" + d.Code).(*action.Context)
			mSec := localCtx.Snap.Finished().Sub(localCtx.Snap.Started()).Nanoseconds() / 1000 / 1000
			if localCtx.Get("success").(bool) {
				rowcount := localCtx.GetDef("rowcount", 0).(int)
				rowsaffected := localCtx.GetDef("rowsaffected", int64(0)).(int64)
				fmt.Printf("%s: Success. Elapsed time:%d msec. Rows count:%d  Rows affected: %d\n", d.Code, mSec, rowcount, rowsaffected)
			} else {
				fmt.Printf("%s: Failed! Elapsed time:%d msec. Error message: %s \n", d.Code, mSec, strings.Replace(localCtx.Snap.Error().Error(), "\n", " ", -1))
			}
		}
	}
	return nil
}
