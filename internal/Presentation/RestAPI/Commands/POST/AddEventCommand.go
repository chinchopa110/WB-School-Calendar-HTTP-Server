package POST

import (
	"WB2/internal/Application/Contracts/UserServices"
	Result2 "WB2/internal/Presentation/RestAPI/Commands/Result"
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

func (c AddEventCommand) Execute() Result2.IExecuteResult {
	res, err := c.service.CreateEvent(c.userId, c.date, c.description, c.key)
	if err != nil {
		result := &Result2.AddUserResult{}
		result.SetError(err)
		return result
	}

	return &Result2.EventResult{Event: res}
}
