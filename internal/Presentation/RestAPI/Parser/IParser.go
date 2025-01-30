package Parser

import (
	"WB2/internal/Presentation/RestAPI/Commands"
	"net/http"
)

type IParser interface {
	AddNext(newChain IParser) IParser
	TryParse(r *http.Request) (Commands.ICommand, error)
}
