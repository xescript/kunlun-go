package utils

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect"
	"github.com/xo/dburl"
)

// DbUrlToDsn convert database url to dsn
func DbUrlToDsn(dbRawUrl string) (retDialect string, dsn string, retErr error) {
	u, err := dburl.Parse(dbRawUrl)
	if err != nil {
		retErr = fmt.Errorf("failed to parse url: %v", err)
		return
	}
	dsn = u.DSN

	switch strings.ToLower(u.URL.Scheme) {
	case "sqlite3", "sqlite", "file", "sq":
		retDialect = dialect.SQLite
		dsn = fmt.Sprintf("file:%s", dsn)
	case "mysql", "my", "mariadb", "maria", "percona", "aurora":
		retDialect = dialect.MySQL
	case "postgres", "pg", "pgsql":
		retDialect = dialect.Postgres
	default:
		retErr = fmt.Errorf("unknown db scheme: %s", u.URL.Scheme)
	}
	return
}
