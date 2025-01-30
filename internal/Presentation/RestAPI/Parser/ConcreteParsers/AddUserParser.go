package ConcreteParsers

import (
	"WB2/internal/Application/Contracts/UserServices"
	"WB2/internal/Presentation/RestAPI/Commands"
	"WB2/internal/Presentation/RestAPI/Commands/POST"
	"WB2/internal/Presentation/RestAPI/Parser"
	"WB2/internal/Presentation/RestAPI/Parser/ConcreteParsers/Validation"
	"log"
	"net/http"
)

type AddUserParser struct {
	Parser.BaseParser
	Service UserServices.IPostService
}

func (p *AddUserParser) TryParse(r *http.Request) (Commands.ICommand, error) {
	if r.URL.Query().Get("type") == "AddUser" {
		log.Printf("Add user command parse %s\n", r.URL.Path)

		key := r.URL.Query().Get("key")

		err := Validation.IsValidUser(key)
		if err != nil {
			return nil, err
		}

		return POST.CreateAddUserCommand(p.Service, key), nil
	}

	return p.BaseParser.TryParse(r)
}
