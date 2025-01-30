package POST

import (
	"WB2/internal/Application/Contracts/UserServices"
	Result2 "WB2/internal/Presentation/RestAPI/Commands/Result"
)

type DeleteEventCommand struct {
	service UserServices.IPostService
	userId  int
	eventId int
	key     string
}

func CreateDeleteEventCommand(service UserServices.IPostService, userId int, key string, eventId int) DeleteEventCommand {
	return DeleteEventCommand{
		service: service,
		userId:  userId,
		eventId: eventId,
		key:     key,
	}
}

func (c DeleteEventCommand) Execute() Result2.IExecuteResult {
	res, err := c.service.DeleteEvent(c.userId, c.eventId, c.key)
	if err != nil {
		result := &Result2.AddUserResult{}
		result.SetError(err)
		return result
	}

	return &Result2.EventResult{Event: res}
}
