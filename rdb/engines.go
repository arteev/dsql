package rdb

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq"               //
	_ "github.com/mattn/go-sqlite3"     //
	_ "github.com/nakagami/firebirdsql" //
)

//KnownEngine - Codes of supported database engines
var KnownEngine = map[string]interface{}{
	"firebirdsql": nil,
	"sqlite3":     nil,
	"postgres":    nil,
}

//Errors
var (
	ErrUnknowEngine = errors.New("An unknown engine of database")
)

//CheckCodeEngine - check supported database engine
func CheckCodeEngine(engine string) {
	if _, ok := KnownEngine[engine]; !ok {
		panic(ErrUnknowEngine)
	}
}

//Open - returns *sql.DB with check of value engine
func Open(name, connectionString string) (*sql.DB, error) {
	CheckCodeEngine(name)	
	return sql.Open(name, connectionString)
}
