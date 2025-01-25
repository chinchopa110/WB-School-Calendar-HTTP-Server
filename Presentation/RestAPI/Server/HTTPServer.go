package Server

import (
	"WB2/Application/Abstractions/Repos"
	"WB2/Application/Services/UserService"
	"WB2/Presentation/RestAPI/Parser"
	"WB2/Presentation/RestAPI/Parser/Factory"
	"encoding/json"
	"log"
	"net/http"
)

type HTTPServer struct {
	parser Parser.IParser
}

func NewHTTPServer(userEventsRepo Repos.IUserEventsRepo) *HTTPServer {
	getService := UserService.CreateGetService(userEventsRepo)
	postService := UserService.CreatePostService(userEventsRepo)

	return &HTTPServer{
		parser: Factory.CreateParser(getService, postService),
	}
}

func (s *HTTPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Start handling request: %s %s", r.Method, r.URL.Path)

	if s.parser == nil {
		log.Println("Error: Parser is not initialized")
		http.Error(w, "Parser is not initialized", http.StatusInternalServerError)
		return
	}

	command, err := s.parser.TryParse(r)
	if err != nil {
		log.Printf("Error parsing request: %s\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := command.Execute()
	if res == nil {
		log.Println("Error: Response is nil")
		http.Error(w, "Response is nil", http.StatusInternalServerError)
		return
	} else if res.GetError() != nil {
		log.Printf("Error executing command: %s\n", res.GetError())
		http.Error(w, res.GetError().Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Error encoding json: %s\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("End handling request: %s %s", r.Method, r.URL.Path)
}
