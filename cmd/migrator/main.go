package main

import (
	"context"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"kunlun/ent"
	"kunlun/internal/utils"
)

// Create schema on database (for development & testing)

func main() {
	dbUrl, ok := os.LookupEnv("DB_URL")
	if !ok {
		log.Println(`Env "DB_URL" not set`)
		os.Exit(1)
	}

	dialect, dsn, err := utils.DbUrlToDsn(dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	client, err := ent.Open(dialect, dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatal(err)
	}
}
