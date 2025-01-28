package main

import (
	"WB2/Application/Services/UserService"
	"WB2/Infrastucture/DataAccess/Repositories"
	"WB2/Presentation/RestAPI/Server"
	"WB2/Presentation/UI/Authentication"
	"WB2/Presentation/UI/Authorized/Menu"
	"WB2/Presentation/middleware"
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

	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Could not close database connection: %s\n", err)
		}
	}()

	userEventsRepo := Repos.NewUserEventsRepo(db)

	httpServer := Server.NewHTTPServer(userEventsRepo)

	var handlerAPI http.Handler = http.HandlerFunc(httpServer.ServeHTTP)
	handlerAPI = middleware.Logging(handlerAPI)
	handlerAPI = middleware.PanicRecovery(handlerAPI)

	authService := Authentication.CreateAuthService(UserService.CreateGetService(userEventsRepo))
	var handlerUI http.Handler = http.HandlerFunc(authService.Handle)
	handlerUI = middleware.Logging(handlerUI)
	handlerUI = middleware.PanicRecovery(handlerUI)

	actionListService := Menu.CreateActionListService(
		UserService.CreateGetService(userEventsRepo),
		UserService.CreatePostService(userEventsRepo))
	var handlerMenuUI http.Handler = http.HandlerFunc(actionListService.Handle)
	handlerMenuUI = middleware.Logging(handlerMenuUI)
	handlerMenuUI = middleware.PanicRecovery(handlerMenuUI)

	mux := http.NewServeMux()

	mux.Handle("/api/", handlerAPI)
	mux.Handle("/ui/", handlerUI)
	mux.Handle("/authorized", handlerMenuUI)
	mux.Handle("/events/day", handlerMenuUI)
	mux.Handle("/events/week", handlerMenuUI)
	mux.Handle("/events/month", handlerMenuUI)
	mux.Handle("/add-event", handlerMenuUI)
	mux.Handle("/update-date", handlerMenuUI)
	mux.Handle("/update-description", handlerMenuUI)
	mux.Handle("/delete-event", handlerMenuUI)

	log.Println("Starting application on :8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
