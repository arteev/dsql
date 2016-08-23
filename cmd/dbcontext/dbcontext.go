package dbcontext

import "database/sql"

//A DBContext connection for repository
type DBContext interface {
	Connection() *sql.DB
	Close() error
	IsCreated() bool
}
