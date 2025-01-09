package POST

import (
	"WB2/Application/Contracts/UserServices"
	"WB2/Presentation/RestAPI/Commands/Result"
)

type UpdateEventDescriptionCommand struct {
	service     UserServices.IPostService
	userId      int
	eventId     int
	description string
	key         string
}

func CreateUpdateEventDescriptionCommand(service UserServices.IPostService, userId int, key string, eventId int, description string) UpdateEventDescriptionCommand {
	return UpdateEventDescriptionCommand{
		service:     service,
		userId:      userId,
		eventId:     eventId,
		description: description,
		key:         key,
	}
}

func (c UpdateEventDescriptionCommand) Execute() Result.IExecuteResult {
	res, err := c.service.UpdateEventDescription(c.userId, c.eventId, c.description, c.key)
	if err != nil {
		result := &Result.AddUserResult{}
		result.SetError(err)
		return result
	}

	return &Result.EventResult{Event: res}
}
