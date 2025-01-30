package Parser

import (
	"WB2/internal/Presentation/RestAPI/Commands"
	"errors"
	"net/http"
)

type BaseParser struct {
	next IParser
}

func (bp *BaseParser) AddNext(newChain IParser) IParser {
	if bp.next != nil {
		bp.next.AddNext(newChain)
	} else {
		bp.next = newChain
	}
	return bp
}

func (bp *BaseParser) TryParse(r *http.Request) (Commands.ICommand, error) {
	if bp.next != nil {
		return bp.next.TryParse(r)
	}
	return nil, errors.New("No parser found ")
}
