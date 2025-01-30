package GET

import (
	"WB2/internal/Application/Contracts/UserServices"
	Result2 "WB2/internal/Presentation/RestAPI/Commands/Result"
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

func (c ForWeekCommand) Execute() Result2.IExecuteResult {
	res, err := c.service.EventsForWeek(c.userId, c.key)
	if err != nil {
		result := &Result2.GetExecuteResult{}
		result.SetError(err)
		return result
	}

	return &Result2.GetExecuteResult{
		Events: res,
	}
}
