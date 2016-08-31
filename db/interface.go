package db

import "errors"

//Errors
var (
	ErrNotFound = errors.New("Database tot found")
)

//A RepositoryDB repository for entities of the databases
type RepositoryDB interface {
	Add(db Database) error
	Update(db Database) error
	Delete(db Database) error
	Refresh(db *Database) error
	FindByCode(code string) (Database, error)
	Close() error
	All() (CollectionRepositoryDB, error)
	AddTags(db *Database, tags ...string) (int,error)
	RemoveTags(db *Database, tags ...string) (int,error)
}

//A CollectionRepositoryDB returns collection of the parameters
type CollectionRepositoryDB interface {
	Get() []Database
	AddFilterEnabled() CollectionRepositoryDB
	AddFilterIncludeDB(code ...string) CollectionRepositoryDB
	AddFilterExcludeDB(code ...string) CollectionRepositoryDB
	AddFilterIncludeEngine(engine ...string) CollectionRepositoryDB
	AddFilterTag(tag ...string) CollectionRepositoryDB
	AddFilterExcludeTag(tag ...string) CollectionRepositoryDB
}
