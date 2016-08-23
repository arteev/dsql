package action

import (
	"database/sql"

	"github.com/arteev/dsql/cmd/db"
	"github.com/arteev/dsql/cmd/rdb/sqlcommand"
)

//Actioner - the function of specific actions for each database
type Actioner func(dbs db.Database, d *sql.DB, cmd *sqlcommand.SQLCommand, ctx *Context) error
