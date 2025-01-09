package GET

import (
	"WB2/Application/Contracts/UserServices"
	"WB2/Presentation/RestAPI/Commands/Result"
)

type ForDayCommand struct {
	service UserServices.IGetService
	userId  int
	key     string
}

func CreateForDayCommand(service UserServices.IGetService, userId int, key string) ForDayCommand {
	return ForDayCommand{
		service: service,
		userId:  userId,
		key:     key,
	}
}

func (c ForDayCommand) Execute() Result.IExecuteResult {
	res, err := c.service.EventsForDay(c.userId, c.key)
	if err != nil {
		result := &Result.GetExecuteResult{}
		result.SetError(err)
		return result
	}

	return &Result.GetExecuteResult{
		Events: res,
	}
}
