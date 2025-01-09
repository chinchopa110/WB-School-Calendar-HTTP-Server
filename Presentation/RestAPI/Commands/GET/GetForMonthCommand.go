package GET

import (
	"WB2/Application/Contracts/UserServices"
	"WB2/Presentation/RestAPI/Commands/Result"
)

type ForMonthCommand struct {
	service UserServices.IGetService
	userId  int
	key     string
}

func CreateForMonthCommand(service UserServices.IGetService, userId int, key string) ForMonthCommand {
	return ForMonthCommand{
		service: service,
		userId:  userId,
		key:     key,
	}
}

func (c ForMonthCommand) Execute() Result.IExecuteResult {
	res, err := c.service.EventsForMonth(c.userId, c.key)
	if err != nil {
		result := &Result.GetExecuteResult{}
		result.SetError(err)
		return result
	}

	return &Result.GetExecuteResult{
		Events: res,
	}
}
