package ConcreteParsers

import (
	"WB2/Application/Contracts/UserServices"
	"WB2/Presentation/RestAPI/Commands"
	"WB2/Presentation/RestAPI/Commands/POST"
	"WB2/Presentation/RestAPI/Parser"
	"net/http"
)

type AddUserParser struct {
	Parser.BaseParser
	Service UserServices.IPostService
}

func (p *AddUserParser) TryParse(r *http.Request) (Commands.ICommand, error) {
	if r.URL.Query().Get("type") == "AddUser" {
		key := r.URL.Query().Get("key")

		return POST.CreateAddUserCommand(p.Service, key), nil
	}

	return p.BaseParser.TryParse(r)
}
