package db

import (
	"bytes"

	"github.com/arteev/logger"
)

type sqliteCollection struct {
	repository            *sQLiteRepository
	filterEnabled         bool
	filerIncludeDatabases []string
	filerIncludeEngine    []string
	filerIncludeTags      []string
}

//getTags load and  append tags for database
func (c *sqliteCollection) getTags(d *Database) {
	sqlStmt := `select id,tag from tags where iddb=?`
	rows, err := c.repository.Connection().Query(sqlStmt, d.ID)
	if err != nil {
		panic(err)
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}()
	for rows.Next() {
		var (
			id  int
			tag string
		)
		if err := rows.Scan(&id, &tag); err != nil {
			panic(err)
		}
		d.Tags = append(d.Tags,
			&Tag{
				ID:    id,
				Value: tag,
				DB:    d,
			})
	}
}

func (c *sqliteCollection) addFilter(data []string) (string, bool) {
	count := len(data)
	if count > 0 {
		var sbuf bytes.Buffer
		strtobuf := func(s string) {
			_, err := sbuf.WriteString(s)
			if err != nil {
				panic(err)
			}
		}
		for i, d := range data {
			if i < count-1 {
				strtobuf("\"" + d + "\",")
			} else {
				strtobuf("\"" + d + "\"")
			}
		}
		return sbuf.String(), true
	}
	return "", false
}

func (c *sqliteCollection) createSQL() string {
	where := false
	var sqlStmt bytes.Buffer

	strtobuf := func(s string) {
		_, err := sqlStmt.WriteString(s)
		if err != nil {
			panic(err)
		}
	}
	And := func() {
		if !where {
			strtobuf("\nwhere")
			where = true
		} else {
			strtobuf("\nand")
		}
	}
	strtobuf(`select id,code,connectionstring,enabled,engine from databases`)
	if c.filterEnabled {
		strtobuf("\nwhere enabled=1")
		where = true
	}
	if dbNames, ok := c.addFilter(c.filerIncludeDatabases); ok {
		And()
		strtobuf("\ncode in (" + dbNames + ")")
	}
	if engines, ok := c.addFilter(c.filerIncludeEngine); ok {
		And()
		strtobuf("\nengine in (" + engines + ")")
	}
	if tags, ok := c.addFilter(c.filerIncludeTags); ok {
		And()
		strtobuf("\nexists( select 1 from tags where  tags.iddb = databases.id and  UPPER(tag) in (" + tags + "))")
	}
	return sqlStmt.String()
}

func (c *sqliteCollection) Get() (res []Database) {

	sqlStmt := c.createSQL()
	logger.Trace.Printf("get sql: %q \n", sqlStmt)

	rows, err := c.repository.Connection().Query(sqlStmt)
	if err != nil {
		panic(err)
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}()

	for rows.Next() {
		var (
			id                int
			code, uri, engine string
			enabled           int
		)
		if err := rows.Scan(&id, &code, &uri, &enabled, &engine); err != nil {
			panic(err)
		}
		d := &Database{
			ID:               id,
			Code:             code,
			ConnectionString: uri,
			Enabled:          enabled != 0,
			Engine:           engine,
		}
		c.getTags(d)
		res = append(res, *d)
	}
	return res
}

func (c *sqliteCollection) AddFilterEnabled() CollectionRepositoryDB {
	c.filterEnabled = true
	return c
}

func (c *sqliteCollection) AddFilterIncludeDB(code ...string) CollectionRepositoryDB {
	//todo: check exists by code
	if len(code) > 0 {
		c.filerIncludeDatabases = append(c.filerIncludeDatabases, code...)
	}
	return c
}

func (c *sqliteCollection) AddFilterIncludeEngine(engine ...string) CollectionRepositoryDB {
	//todo: check exists by engine
	if len(engine) > 0 {
		c.filerIncludeEngine = append(c.filerIncludeEngine, engine...)
	}
	return c
}

func (c *sqliteCollection) AddFilterTag(tag ...string) CollectionRepositoryDB {
	if len(tag) > 0 {
		c.filerIncludeTags = append(c.filerIncludeTags, tag...)
	}
	return c
}
