package db

import "errors"

//Errors
var (
	ErrNotFound = errors.New("Not found")
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
	AddTags(db *Database, tags ...string) error
	RemoveTags(db *Database, tags ...string) error
}

//A CollectionRepositoryDB returns collection of the parameters
type CollectionRepositoryDB interface {
	Get() []Database
	//TODO: refactor filter args string...
	AddFilterEnabled() CollectionRepositoryDB
	AddFilterIncludeDB(code ...string) CollectionRepositoryDB
	AddFilterIncludeEngine(engine ...string) CollectionRepositoryDB

	AddFilterTag(tag ...string) CollectionRepositoryDB
}
