package Migration

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func Initial() {
	connStr := "user=postgres dbname=wb2 password=123 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	err = runMigrations(db)
	if err != nil {
		log.Fatalf("Ошибка при выполнении миграции: %v", err)
	}

	fmt.Println("Миграция выполнена успешно!")
}

func runMigrations(db *sql.DB) error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS Events (
			Id SERIAL PRIMARY KEY,
			Date VARCHAR(10) NOT NULL,
			Description TEXT NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS Users (
			Id SERIAL PRIMARY KEY,
			Key VARCHAR(255) NOT NULL UNIQUE
		);`,
		`CREATE TABLE IF NOT EXISTS UserEvents (
			UserId INT NOT NULL,
			EventId INT NOT NULL,
			FOREIGN KEY (UserId) REFERENCES Users(Id) ON DELETE CASCADE,
			FOREIGN KEY (EventId) REFERENCES Events(Id) ON DELETE CASCADE,
			PRIMARY KEY (UserId, EventId)
		);`,
	}

	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			return fmt.Errorf("ошибка выполнения запроса: %v", err)
		}
	}

	return nil
}
