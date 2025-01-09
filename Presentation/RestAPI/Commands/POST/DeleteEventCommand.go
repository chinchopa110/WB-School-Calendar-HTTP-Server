package POST

import (
	"WB2/Application/Contracts/UserServices"
	"WB2/Presentation/RestAPI/Commands/Result"
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

func (c DeleteEventCommand) Execute() Result.IExecuteResult {
	res, err := c.service.DeleteEvent(c.userId, c.eventId, c.key)
	if err != nil {
		result := &Result.AddUserResult{}
		result.SetError(err)
		return result
	}

	return &Result.EventResult{Event: res}
}
