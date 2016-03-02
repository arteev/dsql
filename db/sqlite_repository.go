package db

import (
	"sync"

	"github.com/arteev/dsql/dbcontext"
	"github.com/arteev/dsql/dbcontextsqlite"
	"github.com/arteev/tern"
)

type sQLiteRepository struct {
	dbcontext.DBContext
}

var (
	once     sync.Once
	instance *sQLiteRepository
)

//GetInstance returns repository for databases
func GetInstance() RepositoryDB {
	once.Do(func() {
		instance = &sQLiteRepository{}
		instance.DBContext = dbcontextsqlite.GetInstance()
	})
	return instance
}

//All. Get all of the database items from database
func (r *sQLiteRepository) All() (CollectionRepositoryDB, error) {
	return &sqliteCollection{
		repository: r,
	}, nil
}

func (r *sQLiteRepository) Add(db Database) error {
	sqlStmt := `insert into databases(code,connectionstring,enabled,engine) values(?,?,?,?)`
	iEnabled := 0
	if db.Enabled {
		iEnabled = 1
	}
	_, err := r.Connection().Exec(sqlStmt, db.Code, db.ConnectionString, iEnabled, db.Engine)
	return err
}

func (r *sQLiteRepository) FindByCode(aCode string) (Database, error) {
	col, err := r.All()
	if err != nil {
		return Database{}, err
	}
	dbs := col.AddFilterIncludeDB(aCode).Get()
	if len(dbs) == 0 {
		err = ErrNotFound
		return Database{}, err
	}
	return dbs[0], nil
}

func (r *sQLiteRepository) Delete(db Database) error {
	res, err := r.Connection().Exec("delete from databases where code=?", db.Code)
	if err != nil {
		return nil
	}
	if a, err := res.RowsAffected(); err != nil {
		return nil
	} else if a == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *sQLiteRepository) Update(db Database) error {
	res, err := r.Connection().Exec("update databases set code=?,connectionstring=?,enabled=?,engine=? where id=?",
		db.Code, db.ConnectionString, tern.Op(db.Enabled, 1, 0), db.Engine, db.ID)
	if err != nil {
		return nil
	}
	if a, err := res.RowsAffected(); err != nil {
		return nil
	} else if a == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *sQLiteRepository) Refresh(db *Database) error {
	d, err := r.FindByCode(db.Code)
	if err != nil {
		return err
	}
	db.Code = d.Code
	db.ConnectionString = d.ConnectionString
	db.Enabled = d.Enabled
	db.ID = d.ID
	db.Tags = nil
	for _, t := range d.Tags {
		db.Tags = append(db.Tags, &Tag{
			ID:    t.ID,
			Value: t.Value,
			DB:    db,
		})
	}
	return nil
}

func (r *sQLiteRepository) AddTags(db *Database, tags ...string) (int, error) {
	if len(tags) == 0 {
		return 0, nil
	}
	sqlStmt := `insert into tags(iddb,tag) values(?,?)`
	count := 0
	for _, t := range tags {
		if !db.TagExists(t) {
			ex, err := r.Connection().Exec(sqlStmt, db.ID, t)
			if err != nil {
				return 0, err
			}
			if c, err := ex.RowsAffected(); err == nil {
				count += int(c)
			}
		}
	}
	return count, r.Refresh(db)
}

func (r *sQLiteRepository) RemoveTags(db *Database, tags ...string) (int, error) {
	if len(tags) == 0 {
		return 0, nil
	}
	sqlStmt := `delete from tags where iddb=? and Upper(tag)=Upper(?)`
	count := 0
	for _, t := range tags {
		if db.TagExists(t) {
			ex, err := r.Connection().Exec(sqlStmt, db.ID, t)
			if err != nil {
				return 0, err
			}
			if c, err := ex.RowsAffected(); err == nil {
				count += int(c)
			}
		}
	}
	return count, r.Refresh(db)
}
