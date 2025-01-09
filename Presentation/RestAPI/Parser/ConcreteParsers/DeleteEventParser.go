package ConcreteParsers

import (
	"WB2/Application/Contracts/UserServices"
	"WB2/Presentation/RestAPI/Commands"
	"WB2/Presentation/RestAPI/Commands/POST"
	"WB2/Presentation/RestAPI/Parser"
	"net/http"
	"strconv"
)

type DeleteEventParser struct {
	Parser.BaseParser
	Service UserServices.IPostService
}

func (p *DeleteEventParser) TryParse(r *http.Request) (Commands.ICommand, error) {
	if r.URL.Query().Get("type") == "DeleteEvent" {
		userIDStr := r.URL.Query().Get("user_id")
		key := r.URL.Query().Get("key")
		eventIDStr := r.URL.Query().Get("event_id")

		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			return nil, err
		}

		eventID, err := strconv.Atoi(eventIDStr)
		if err != nil {
			return nil, err
		}

		return POST.CreateDeleteEventCommand(p.Service, userID, key, eventID), nil
	}

	return p.BaseParser.TryParse(r)
}
