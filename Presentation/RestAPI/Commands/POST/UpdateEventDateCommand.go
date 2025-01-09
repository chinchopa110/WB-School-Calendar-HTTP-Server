package POST

import (
	"WB2/Application/Contracts/UserServices"
	"WB2/Presentation/RestAPI/Commands/Result"
)

type UpdateEventDateCommand struct {
	service UserServices.IPostService
	userId  int
	eventId int
	date    string
	key     string
}

func CreateUpdateEventDateCommand(service UserServices.IPostService, userId int, key string, eventId int, date string) UpdateEventDateCommand {
	return UpdateEventDateCommand{
		service: service,
		userId:  userId,
		eventId: eventId,
		date:    date,
		key:     key,
	}
}

func (c UpdateEventDateCommand) Execute() Result.IExecuteResult {
	res, err := c.service.UpdateEventDate(c.userId, c.eventId, c.date, c.key)
	if err != nil {
		result := &Result.AddUserResult{}
		result.SetError(err)
		return result
	}

	return &Result.EventResult{Event: res}
}
