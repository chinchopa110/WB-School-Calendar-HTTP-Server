package ConcreteParsers

import (
	"WB2/Application/Contracts/UserServices"
	"WB2/Presentation/RestAPI/Commands"
	"WB2/Presentation/RestAPI/Commands/GET"
	"WB2/Presentation/RestAPI/Parser"
	"log"
	"net/http"
	"strconv"
)

type GetForMonthParser struct {
	Parser.BaseParser
	Service UserServices.IGetService
}

func (p *GetForMonthParser) TryParse(r *http.Request) (Commands.ICommand, error) {
	if r.URL.Query().Get("type") == "GetForMonth" {
		userIDStr := r.URL.Query().Get("user_id")
		key := r.URL.Query().Get("key")

		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			return nil, err
		}

		log.Printf("Get events for month command parse %s\n", r.URL.Path)

		return GET.CreateForMonthCommand(p.Service, userID, key), nil
	}

	return p.BaseParser.TryParse(r)
}
