package Result

import (
	"WB2/Application/Domain"
)

type EventResult struct {
	error error
	Event Domain.Event
}

func (t *EventResult) SetError(Error error) {
	t.error = Error
}

func (t *EventResult) GetError() error {
	return t.error
}
