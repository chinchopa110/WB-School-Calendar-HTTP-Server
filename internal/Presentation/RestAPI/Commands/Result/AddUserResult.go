package Result

import (
	"WB2/internal/Application/Domain"
)

type AddUserResult struct {
	error error
	User  *Domain.User
}

func (t *AddUserResult) SetError(Error error) {
	t.error = Error
}

func (t *AddUserResult) GetError() error {
	return t.error
}
