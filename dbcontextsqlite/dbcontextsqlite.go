package dbcontextsqlite

import (
	"database/sql"
	"sync"

	"github.com/arteev/dsql/dbcontext"
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

//RepositoryFile - current file of repository
var RepositoryFile string

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
	if RepositoryFile == "" {
		RepositoryFile = "./repository.sqlite"
	}
	r.dbs, err = sql.Open("sqlite3", RepositoryFile)
	if err != nil {
		return err
	}
	if !r.isExistsMetadata() {
		if err := r.createDataBase(); err != nil {
			return err
		}
	}
	return nil
}

func (r *DBContextSQLite) isExistsMetadata() bool {
	rw := r.dbs.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='databases';")
	var name string
	err := rw.Scan(&name)
	return err == nil && name == "databases"
}

//createDataBase Create new repository with metadata DB
func (r *DBContextSQLite) createDataBase() error {
	sqlStmt := `CREATE TABLE databases (
    id               INTEGER NOT NULL
                             PRIMARY KEY AUTOINCREMENT
                             UNIQUE,
    code             TEXT    NOT NULL
                             UNIQUE,
    connectionstring TEXT    NOT NULL,
    enabled          INTEGER NOT NULL
                             DEFAULT 1
                             CHECK (enabled IN (0,1),
    engine           TEXT    NOT NULL)
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
