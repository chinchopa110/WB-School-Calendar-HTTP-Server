package ConcreteParsers

import (
	"WB2/internal/Application/Contracts/UserServices"
	"WB2/internal/Presentation/RestAPI/Commands"
	"WB2/internal/Presentation/RestAPI/Commands/POST"
	"WB2/internal/Presentation/RestAPI/Parser"
	"WB2/internal/Presentation/RestAPI/Parser/ConcreteParsers/Validation"
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
		log.Printf("Update event date command parse %s\n", r.URL.Path)

		userIDStr := r.URL.Query().Get("user_id")
		key := r.URL.Query().Get("key")
		eventIDStr := r.URL.Query().Get("event_id")
		newDateStr := r.URL.Query().Get("new_date")

		err := Validation.IsValidDate(newDateStr)
		if err != nil {
			return nil, err
		}

		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			return nil, err
		}

		eventID, err := strconv.Atoi(eventIDStr)
		if err != nil {
			return nil, err
		}

		return POST.CreateUpdateEventDateCommand(p.Service, userID, key, eventID, newDateStr), nil
	}

	return p.BaseParser.TryParse(r)
}
