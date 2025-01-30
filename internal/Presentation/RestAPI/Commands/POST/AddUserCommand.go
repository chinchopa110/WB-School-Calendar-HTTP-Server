package POST

import (
	"WB2/internal/Application/Contracts/UserServices"
	"WB2/internal/Application/Domain"
	Result2 "WB2/internal/Presentation/RestAPI/Commands/Result"
)

type AddUserCommand struct {
	service UserServices.IPostService
	user    *Domain.User
}

func CreateAddUserCommand(service UserServices.IPostService, key string) AddUserCommand {
	user := Domain.User{Key: key}
	return AddUserCommand{
		user:    &user,
		service: service}
}

func (c AddUserCommand) Execute() Result2.IExecuteResult {
	res, err := c.service.AddUser(c.user)
	if err != nil {
		result := &Result2.AddUserResult{}
		result.SetError(err)
		return result
	}

	return &Result2.AddUserResult{User: res}
}
