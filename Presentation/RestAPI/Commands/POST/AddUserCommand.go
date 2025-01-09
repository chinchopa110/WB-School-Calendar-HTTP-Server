package POST

import (
	"WB2/Application/Contracts/UserServices"
	"WB2/Application/Domain"
	"WB2/Presentation/RestAPI/Commands/Result"
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

func (c AddUserCommand) Execute() Result.IExecuteResult {
	res, err := c.service.AddUser(c.user)
	if err != nil {
		result := &Result.AddUserResult{}
		result.SetError(err)
		return result
	}

	return &Result.AddUserResult{User: res}
}
