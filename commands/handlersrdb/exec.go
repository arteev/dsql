package handlersrdb

import (
	"database/sql"

	"github.com/arteev/dsql/db"
	"github.com/arteev/dsql/rdb/action"
	"github.com/arteev/dsql/rdb/sqlcommand"
	"github.com/arteev/logger"
)

//Exec - execute sql for the databases
func Exec(dbContext db.Database, dbHandle *sql.DB, cmd *sqlcommand.SQLCommand, ctx *action.Context) error {
	logger.Trace.Println("run exec", dbContext.Code)
	defer logger.Trace.Println("done exec", dbContext.Code)

	localCtx := ctx.Get("context" + dbContext.Code).(*action.Context)
	var pint []interface{}
	for _, p := range cmd.Params {
		pint = append(pint, p)
	}

	res, err := dbHandle.Exec(cmd.Script, pint...)
	if err != nil {
		return err
	}
	ra, err := res.RowsAffected()
	if err != nil {
		return err
	}
	localCtx.IncInt64("rowsaffected", ra)
	ctx.IncInt64("rowsaffected", ra)
	return nil
}
