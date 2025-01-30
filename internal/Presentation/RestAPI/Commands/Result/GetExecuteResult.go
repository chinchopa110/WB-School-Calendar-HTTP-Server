package Result

import (
	"WB2/internal/Application/Domain"
)

type GetExecuteResult struct {
	error  error
	Events []Domain.Event
}

func (t *GetExecuteResult) SetError(Error error) {
	t.error = Error
}

func (t *GetExecuteResult) GetError() error {
	return t.error
}
