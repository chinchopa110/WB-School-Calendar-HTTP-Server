package main

import (
	"WB2/Infrastucture/DataAccess/Repositories"
	"WB2/Presentation/RestAPI/Server"
	"WB2/Presentation/RestAPI/middleware"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=postgres password=123 dbname=wb2 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Could not connect to the database: %s\n", err)
	}
	defer db.Close()

	userEventsRepo := Repos.NewUserEventsRepo(db)

	httpServer := Server.NewHTTPServer(userEventsRepo)

	var handler http.Handler = http.HandlerFunc(httpServer.ServeHTTP)

	handler = middleware.Logging(handler)
	handler = middleware.PanicRecovery(handler)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
