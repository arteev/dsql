package rdb

import (
	"database/sql"
	"fmt"

	_ "github.com/arteev/firebirdsql" //
	//	_ "github.com/nakagami/firebirdsql" //
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3" //
)

//KnownEngine - Codes of supported database engines
var KnownEngine = [...]string{
	"firebirdsql",
	"sqlite3",
	"postgres",
}

//CheckCodeEngine - check supported database engine
func CheckCodeEngine(eng string) {
	for _, e := range KnownEngine {
		if e == eng {
			return
		}
	}
	panic(fmt.Errorf("Unknow engine %q", eng))
}

//Open - returns *sql.DB with check of value engine
func Open(name, connectionString string) (*sql.DB, error) {
	CheckCodeEngine(name)
	return sql.Open(name, connectionString)
}
