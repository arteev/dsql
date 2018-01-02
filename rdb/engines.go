package rdb

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq"               //
	_ "github.com/mattn/go-sqlite3"     //
	_ "github.com/nakagami/firebirdsql" //
)

//KnownEngine - Codes of supported database engines
var KnownEngine = []string{
	"firebirdsql", "sqlite3", "postgres"}

//Errors
var (
	ErrUnknowEngine = errors.New("Unknown engine of database")
)

//CheckCodeEngine - check supported database engine
func CheckCodeEngine(engine string) error {
	for _, e := range KnownEngine {
		if e == engine {
			return nil
		}
	}
	return ErrUnknowEngine
}

//Open - returns *sql.DB with check of value engine
func Open(name, connectionString string) (*sql.DB, error) {
	if err := CheckCodeEngine(name); err != nil {
		return nil, err
	}
	return sql.Open(name, connectionString)
}
