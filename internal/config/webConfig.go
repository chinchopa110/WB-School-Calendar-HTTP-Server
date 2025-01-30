package config

import (
	UserService2 "WB2/internal/Application/Services/UserService"
	Repos "WB2/internal/Infrastucture/DataAccess/Repositories"
	"WB2/internal/Presentation/RestAPI/Server"
	"WB2/internal/Presentation/UI/Authentication"
	"WB2/internal/Presentation/UI/Authorized/Menu"
	"WB2/internal/Presentation/middleware"
	"log"
	"net/http"
)

func GetUpServer(userEventsRepo *Repos.UserEventsRepo) {
	httpServer := Server.NewHTTPServer(userEventsRepo)

	var handlerAPI http.Handler = http.HandlerFunc(httpServer.ServeHTTP)
	handlerAPI = middleware.Logging(handlerAPI)
	handlerAPI = middleware.PanicRecovery(handlerAPI)

	authService := Authentication.CreateAuthService(UserService2.CreateGetService(userEventsRepo))
	var handlerUI http.Handler = http.HandlerFunc(authService.Handle)
	handlerUI = middleware.Logging(handlerUI)
	handlerUI = middleware.PanicRecovery(handlerUI)

	actionListService := Menu.CreateActionListService(
		UserService2.CreateGetService(userEventsRepo),
		UserService2.CreatePostService(userEventsRepo))
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
