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

type AddEventParser struct {
	Parser.BaseParser
	Service UserServices.IPostService
}

func (p *AddEventParser) TryParse(r *http.Request) (Commands.ICommand, error) {
	if r.URL.Query().Get("type") == "AddEvent" {
		userIDStr := r.URL.Query().Get("user_id")
		key := r.URL.Query().Get("key")
		date := r.URL.Query().Get("date")
		description := r.URL.Query().Get("description")

		log.Printf("Add event command parse %s\n", r.URL.Path)

		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			return nil, err
		}

		return POST.CreateAddEventCommand(p.Service, userID, key, date, description), nil
	}

	return p.BaseParser.TryParse(r)
}
