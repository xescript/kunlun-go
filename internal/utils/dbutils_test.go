package utils

import (
	"testing"

	"entgo.io/ent/dialect"
	"github.com/stretchr/testify/assert"
)

func TestDBUrlToDsn(t *testing.T) {
	dialect1, dsn1, err := DbUrlToDsn("sqlite://./db.sqlite3?_fk=1&cache=shared")
	assert.Nil(t, err)
	assert.Equal(t, dialect.SQLite, dialect1)
	assert.Equal(t, "file:./db.sqlite3?_fk=1&cache=shared", dsn1)

	dialect2, dsn2, err := DbUrlToDsn("mysql://root:123456@localhost/dbname")
	assert.Nil(t, err)
	assert.Equal(t, dialect.MySQL, dialect2)
	assert.Equal(t, "root:123456@tcp(localhost:3306)/dbname", dsn2)

	dialect3, dsn3, err := DbUrlToDsn("pg://root:123456@localhost/dbname")
	assert.Nil(t, err)
	assert.Equal(t, dialect.Postgres, dialect3)
	assert.Equal(t, "dbname=dbname host=localhost password=123456 user=root", dsn3)
}
