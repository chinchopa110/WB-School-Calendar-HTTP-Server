package ConcreteParsers

import (
	"WB2/Application/Contracts/UserServices"
	"WB2/Presentation/RestAPI/Commands"
	"WB2/Presentation/RestAPI/Commands/POST"
	"WB2/Presentation/RestAPI/Parser"
	"log"
	"net/http"
)

type AddUserParser struct {
	Parser.BaseParser
	Service UserServices.IPostService
}

func (p *AddUserParser) TryParse(r *http.Request) (Commands.ICommand, error) {
	if r.URL.Query().Get("type") == "AddUser" {
		key := r.URL.Query().Get("key")

		log.Printf("Add user command parse %s\n", r.URL.Path)

		return POST.CreateAddUserCommand(p.Service, key), nil
	}

	return p.BaseParser.TryParse(r)
}
