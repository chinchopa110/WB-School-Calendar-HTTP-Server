package GET

import (
	"WB2/Application/Contracts/UserServices"
	"WB2/Presentation/RestAPI/Commands/Result"
)

type ForWeekCommand struct {
	service UserServices.IGetService
	userId  int
	key     string
}

func CreateForWeekCommand(service UserServices.IGetService, userId int, key string) ForWeekCommand {
	return ForWeekCommand{
		service: service,
		userId:  userId,
		key:     key,
	}
}

func (c ForWeekCommand) Execute() Result.IExecuteResult {
	res, err := c.service.EventsForWeek(c.userId, c.key)
	if err != nil {
		result := &Result.GetExecuteResult{}
		result.SetError(err)
		return result
	}

	return &Result.GetExecuteResult{
		Events: res,
	}
}
