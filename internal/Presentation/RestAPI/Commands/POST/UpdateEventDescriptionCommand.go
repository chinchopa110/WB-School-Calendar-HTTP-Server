package POST

import (
	"WB2/internal/Application/Contracts/UserServices"
	Result2 "WB2/internal/Presentation/RestAPI/Commands/Result"
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

func (c UpdateEventDescriptionCommand) Execute() Result2.IExecuteResult {
	res, err := c.service.UpdateEventDescription(c.userId, c.eventId, c.description, c.key)
	if err != nil {
		result := &Result2.AddUserResult{}
		result.SetError(err)
		return result
	}

	return &Result2.EventResult{Event: res}
}
