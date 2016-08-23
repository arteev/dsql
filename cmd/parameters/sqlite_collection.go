package parameters

import (
	"database/sql"

	"github.com/arteev/dsql/cmd/dbcontext"
)

type sqliteCollectionParams struct {
	ctx dbcontext.DBContext
}

func (c *sqliteCollectionParams) Get() (res []Parameter) {

	sqlStmt := `select id,name,value,description from parameters;`
	rows, err := c.ctx.Connection().Query(sqlStmt)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			panic(err)
		}
	}()

	for rows.Next() {
		var (
			id    int
			name  string
			value sql.NullString
			desc  sql.NullString
		)
		if err := rows.Scan(&id, &name, &value, &desc); err != nil {
			panic(err)
		}
		res = append(res,
			Parameter{
				ID:          id,
				Name:        name,
				Value:       value.String,
				Description: desc.String,
			})
	}
	return res
}
