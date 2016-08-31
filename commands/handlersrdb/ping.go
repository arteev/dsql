package handlersrdb

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/arteev/dsql/db"
	"github.com/arteev/dsql/rdb/action"
	"github.com/arteev/dsql/rdb/sqlcommand"
)

//PingHandler - Ping the databases
func PingHandler(d db.Database, dh *sql.DB, cmd *sqlcommand.SQLCommand, ctx *action.Context) error {
	tstart := time.Now()
	err := dh.Ping()
	//TODO: переделать вывод статистики через триггер т.к. snap фиксирует старт и завершение
	mSec := time.Now().Sub(tstart).Nanoseconds() / 1000 / 1000
	if err != nil {
		/*if !ctx.GetDef("silent",true).(bool) {
			fmt.Fprintf(os.Stderr, "%s ping: %v\n", d.Code, strings.TrimSpace(err.Error()))
		}*/
		return err
	}

	fmt.Printf("%s ping: %v msec\n", d.Code, mSec)
	return nil
}
