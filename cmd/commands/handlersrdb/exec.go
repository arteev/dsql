package handlersrdb

import (
	"database/sql"

	"github.com/arteev/dsql/cmd/db"
	"github.com/arteev/dsql/cmd/rdb/action"
	"github.com/arteev/dsql/cmd/rdb/sqlcommand"
	"github.com/arteev/logger"
)

//Exec - execute sql for the databases
func Exec(dbContext db.Database, dbHandle *sql.DB, cmd *sqlcommand.SQLCommand, ctx *action.Context) error {
	logger.Trace.Println("run exec", dbContext.Code)
	defer logger.Trace.Println("done exec", dbContext.Code)

	localCtx := ctx.Get("context" + dbContext.Code).(*action.Context)
	commit := ctx.GetDef("commit",false).(bool)
	var pint []interface{}
	for _, p := range cmd.Params {
		pint = append(pint, p)
	}


	tx,err:=dbHandle.Begin()	
	if err!=nil {
		return err
	}
	sqmt, err := tx.Prepare(cmd.Script)
	if err != nil {
		return err
	}	
	defer sqmt.Close()	 		 
	res, err := sqmt.Exec(pint...)
	if err != nil {
		return err
	}	
	defer sqmt.Close()

	if commit {
		logger.Debug.Printf("%s Transaction commited",dbContext.Code)		
		if err:=tx.Commit();err != nil {
			return err
		}
	} else {	    	
		if err:=tx.Rollback();err != nil {
			return err
		}
		logger.Warn.Printf("%s: Transaction rollback.Use a special flag for commit the transaction\n",dbContext.Code)
	}
	ra, err := res.RowsAffected()
	if err != nil {		
		return err
	} 
	
	localCtx.IncInt64("rowsaffected", ra)
	ctx.IncInt64("rowsaffected", ra)
	return nil
}
