package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func GetUpSQL(connStr string) *sql.DB {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Could not connect to the database: %s\n", err)
	}

	return db
}
