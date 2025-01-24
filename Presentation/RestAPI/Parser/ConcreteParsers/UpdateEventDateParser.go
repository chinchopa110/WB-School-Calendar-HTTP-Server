package ConcreteParsers

import (
	"WB2/Application/Contracts/UserServices"
	"WB2/Presentation/RestAPI/Commands"
	"WB2/Presentation/RestAPI/Commands/POST"
	"WB2/Presentation/RestAPI/Parser"
	"log"
	"net/http"
	"strconv"
)

type UpdateEventDateParser struct {
	Parser.BaseParser
	Service UserServices.IPostService
}

func (p *UpdateEventDateParser) TryParse(r *http.Request) (Commands.ICommand, error) {
	if r.URL.Query().Get("type") == "UpdateEventDate" {
		userIDStr := r.URL.Query().Get("user_id")
		key := r.URL.Query().Get("key")
		eventIDStr := r.URL.Query().Get("event_id")
		newDateStr := r.URL.Query().Get("new_date")

		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			return nil, err
		}

		eventID, err := strconv.Atoi(eventIDStr)
		if err != nil {
			return nil, err
		}

		log.Printf("Update event date command parse %s\n", r.URL.Path)

		return POST.CreateUpdateEventDateCommand(p.Service, userID, key, eventID, newDateStr), nil
	}

	return p.BaseParser.TryParse(r)
}
