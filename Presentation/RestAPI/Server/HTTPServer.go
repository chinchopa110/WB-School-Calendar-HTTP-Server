package Server

import (
	"WB2/Application/Abstractions/Repos"
	"WB2/Application/Services/UserService"
	"WB2/Presentation/RestAPI/Parser"
	"WB2/Presentation/RestAPI/Parser/Factory"
	"encoding/json"
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
	if s.parser == nil {
		http.Error(w, "Parser is not initialized", http.StatusInternalServerError)
		return
	}

	command, err := s.parser.TryParse(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := command.Execute()
	if res.GetError() != nil {
		http.Error(w, res.GetError().Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
