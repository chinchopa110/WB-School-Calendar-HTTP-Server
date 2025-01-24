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

type GetForWeekParser struct {
	Parser.BaseParser
	Service UserServices.IGetService
}

func (p *GetForWeekParser) TryParse(r *http.Request) (Commands.ICommand, error) {
	if r.URL.Query().Get("type") == "GetForWeek" {
		userIDStr := r.URL.Query().Get("user_id")
		key := r.URL.Query().Get("key")

		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			return nil, err
		}

		log.Printf("Get events for week command parse %s\n", r.URL.Path)

		return GET.CreateForWeekCommand(p.Service, userID, key), nil
	}

	return p.BaseParser.TryParse(r)
}
