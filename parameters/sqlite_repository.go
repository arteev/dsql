package parameters

import (
	"errors"
	"strings"
	"sync"

	"github.com/arteev/dsql/dbcontext"
	"github.com/arteev/dsql/dbcontextsqlite"
)

//Error for parameters
var (
	ErrNotFound = errors.New("Not found")
)

type sqliteRepositoryParams struct {
	dbcontext.DBContext
}

var (
	once     sync.Once
	instance *sqliteRepositoryParams
)

func (r *sqliteRepositoryParams) setPreDefinedParams() error {
	for _, param := range definedParams {
		err := r.Add(MustParameter(param.Name, param.Default, param.Description))
		if err != nil {
			return err
		}
	}
	return nil
}

//GetInstance create repository of parameters
func GetInstance() RepositoryParams {
	once.Do(func() {
		instance = &sqliteRepositoryParams{}
		instance.DBContext = dbcontextsqlite.GetInstance()

		//Create prdefined params
		if instance.DBContext.IsCreated() {
			if err := instance.setPreDefinedParams(); err != nil {
				panic(err)
			}
		}
	})
	return instance
}

//All. Get all of the parameters items from repository
func (r *sqliteRepositoryParams) All() (CollectionRepositoryParams, error) {
	return &sqliteCollectionParams{
		ctx: r.DBContext,
	}, nil

}
func (r *sqliteRepositoryParams) Add(p Parameter) error {
	sqlStmt := `insert into parameters(name,value,description) values(?,?,?);`
	var val interface{}
	var desc interface{}
	if p.Value != "" {
		val = p.Value
	}
	if p.Description != "" {
		desc = p.Description
	}
	_, err := r.Connection().Exec(sqlStmt, p.Name, val, desc)
	return err
}

func (r *sqliteRepositoryParams) Update(p Parameter) error {
	var val interface{}
	var desc interface{}
	if p.Value != "" {
		val = p.Value
	}
	if p.Description != "" {
		desc = p.Description
	}
	res, err := r.Connection().Exec("update parameters set value=?, description=? where id=?",
		val, desc, p.ID)
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

func (r *sqliteRepositoryParams) Delete(p Parameter) error {
	sqlStmt := `delete from parameters where id=?;`
	_, err := r.Connection().Exec(sqlStmt, p.ID)
	return err
}

func (r *sqliteRepositoryParams) FindByName(name string) (Parameter, error) {
	params, err := r.All()
	if err != nil {
		return Parameter{}, err
	}
	for _, param := range params.Get() {
		if strings.ToUpper(param.Name) == strings.ToUpper(name) {
			return param, nil
		}
	}
	return Parameter{}, ErrNotFound
}
