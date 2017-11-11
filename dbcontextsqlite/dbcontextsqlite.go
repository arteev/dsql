package dbcontextsqlite

import (
	"database/sql"
	"sync"

	"github.com/arteev/dsql/dbcontext"
	"github.com/arteev/dsql/repository"
)

//DBContextSQLite - impementation DBContext
type DBContextSQLite struct {
	isCreated bool
	dbs       *sql.DB
}

var (
	once     sync.Once
	instance *DBContextSQLite
)

//GetInstance get global DBContext
func GetInstance() dbcontext.DBContext {
	once.Do(func() {
		instance = New()
	})
	return instance
}

//New create DBContextSQLite
func New() *DBContextSQLite {
	result := &DBContextSQLite{}
	if err := result.init(); err != nil {
		panic(err)
	}
	return result
}

func (r *DBContextSQLite) init() error {
	var err error
	if err = repository.PrepareLocation(); err != nil {
		return err
	}

	r.dbs, err = sql.Open("sqlite3", repository.GetRepositoryFile())
	if err != nil {
		return err
	}

	if err := r.migrate(); err != nil {
		return err
	}

	return nil
}

func (r *DBContextSQLite) migrated() bool {
	rw := r.dbs.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='databases';")
	var name string
	err := rw.Scan(&name)
	return err == nil && name == "databases"
}

//migrate Create new repository with metadata DB
func (r *DBContextSQLite) migrate() error {
	if r.migrated() {
		return nil
	}
	sqlStmt := `CREATE TABLE databases (
    id               INTEGER NOT NULL
                             PRIMARY KEY AUTOINCREMENT
                             UNIQUE,
    code             TEXT    NOT NULL
                             UNIQUE,
    connectionstring TEXT    NOT NULL,
    enabled          INTEGER NOT NULL
                             DEFAULT 1
                             CHECK (enabled IN (0,1)),
    engine           TEXT    NOT NULL
);

CREATE TABLE parameters (
	id               INTEGER NOT NULL
                             PRIMARY KEY AUTOINCREMENT
                             UNIQUE,
    name             TEXT    NOT NULL
                             UNIQUE,
	value            TEXT,
	description      TEXT

	);
    
CREATE TABLE tags (
    id   INTEGER PRIMARY KEY AUTOINCREMENT
                 UNIQUE
                 NOT NULL,
    tag  TEXT    NOT NULL,
    iddb INTEGER REFERENCES databases (id) ON DELETE CASCADE
                                           ON UPDATE CASCADE
                 NOT NULL
);


CREATE UNIQUE INDEX IDX_TAGS_DB ON tags (
    tag COLLATE NOCASE,
    iddb
);

    
`
	_, err := r.dbs.Exec(sqlStmt)
	if err != nil {
		return err
	}
	r.isCreated = true
	return nil
}

//Connection returns current connection *sql.DB
func (r *DBContextSQLite) Connection() *sql.DB {
	return r.dbs
}

//Close current conection
func (r *DBContextSQLite) Close() error {
	return r.dbs.Close()
}

//IsCreated returns true if repository created in this sessions
func (r *DBContextSQLite) IsCreated() bool {
	return r.isCreated
}
