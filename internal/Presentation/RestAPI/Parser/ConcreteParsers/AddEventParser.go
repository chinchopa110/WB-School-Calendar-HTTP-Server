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

type AddEventParser struct {
	Parser.BaseParser
	Service UserServices.IPostService
}

func (p *AddEventParser) TryParse(r *http.Request) (Commands.ICommand, error) {
	if r.URL.Query().Get("type") == "AddEvent" {
		log.Printf("Add event command parse %s\n", r.URL.Path)

		userIDStr := r.URL.Query().Get("user_id")
		key := r.URL.Query().Get("key")
		date := r.URL.Query().Get("date")
		description := r.URL.Query().Get("description")

		err := Validation.IsValidEvent(userIDStr, key, date, description)
		if err != nil {
			return nil, err
		}

		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			return nil, err
		}

		return POST.CreateAddEventCommand(p.Service, userID, key, date, description), nil
	}

	return p.BaseParser.TryParse(r)
}
