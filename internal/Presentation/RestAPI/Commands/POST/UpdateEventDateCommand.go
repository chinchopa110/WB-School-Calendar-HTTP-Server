package POST

import (
	"WB2/internal/Application/Contracts/UserServices"
	Result2 "WB2/internal/Presentation/RestAPI/Commands/Result"
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

func (c UpdateEventDateCommand) Execute() Result2.IExecuteResult {
	res, err := c.service.UpdateEventDate(c.userId, c.eventId, c.date, c.key)
	if err != nil {
		result := &Result2.AddUserResult{}
		result.SetError(err)
		return result
	}

	return &Result2.EventResult{Event: res}
}
