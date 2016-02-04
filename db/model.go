package db

import (
	"bytes"
	"strings"
)

//A Tag item
type Tag struct {
	ID    int
	DB    *Database
	Value string
}

//A Database item
type Database struct {
	ID               int
	Code             string
	ConnectionString string
	Enabled          bool
	Engine           string
	Tags             []*Tag
}

//TagsComma returns the tags as string through a separator
func (d Database) TagsComma(sep string) string {
	var buf bytes.Buffer
	for _, t := range d.Tags {
		_, err := buf.WriteString(t.Value + sep)
		if err != nil {
			panic(err)
		}
	}
	return buf.String()
}

//TagExists check tag exists in Database.Tags
func (d *Database) TagExists(value string) bool {
	for _, t := range d.Tags {
		if strings.ToUpper(t.Value) == strings.ToUpper(value) {
			return true
		}
	}
	return false
}
