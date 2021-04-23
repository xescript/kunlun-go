package main

import (
	"context"
	"log"

	"entgo.io/ent/dialect"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"kunlun/ent"
)

// Create schema on database (for development & testing)

func main() {
	client, err := ent.Open(dialect.SQLite, "file:./db.sqlite3?_fk=1")
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatal(err)
	}
}
