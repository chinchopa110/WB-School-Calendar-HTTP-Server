package GET

import (
	"WB2/internal/Application/Contracts/UserServices"
	Result2 "WB2/internal/Presentation/RestAPI/Commands/Result"
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

func (c ForMonthCommand) Execute() Result2.IExecuteResult {
	res, err := c.service.EventsForMonth(c.userId, c.key)
	if err != nil {
		result := &Result2.GetExecuteResult{}
		result.SetError(err)
		return result
	}

	return &Result2.GetExecuteResult{
		Events: res,
	}
}
