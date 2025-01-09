package ConcreteParsers

import (
	"WB2/Application/Contracts/UserServices"
	"WB2/Presentation/RestAPI/Commands"
	"WB2/Presentation/RestAPI/Commands/GET"
	"WB2/Presentation/RestAPI/Parser"
	"net/http"
	"strconv"
)

type GetForDayParser struct {
	Parser.BaseParser
	Service UserServices.IGetService
}

func (p *GetForDayParser) TryParse(r *http.Request) (Commands.ICommand, error) {
	if r.URL.Query().Get("type") == "GetForDay" {
		userIDStr := r.URL.Query().Get("user_id")
		key := r.URL.Query().Get("key")

		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			return nil, err
		}

		return GET.CreateForDayCommand(p.Service, userID, key), nil
	}

	return p.BaseParser.TryParse(r)
}
