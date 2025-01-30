package Migration

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

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
