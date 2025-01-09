package POST

import (
	"WB2/Application/Contracts/UserServices"
	"WB2/Presentation/RestAPI/Commands/Result"
)

type AddEventCommand struct {
	service     UserServices.IPostService
	userId      int
	date        string
	description string
	key         string
}

func CreateAddEventCommand(service UserServices.IPostService, userId int, key string, date string, description string) AddEventCommand {
	return AddEventCommand{
		service:     service,
		userId:      userId,
		date:        date,
		description: description,
		key:         key,
	}
}

func (c AddEventCommand) Execute() Result.IExecuteResult {
	res, err := c.service.CreateEvent(c.userId, c.date, c.description, c.key)
	if err != nil {
		result := &Result.AddUserResult{}
		result.SetError(err)
		return result
	}

	return &Result.EventResult{Event: res}
}
