package GET

import (
	"WB2/internal/Application/Contracts/UserServices"
	Result2 "WB2/internal/Presentation/RestAPI/Commands/Result"
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

func (c ForDayCommand) Execute() Result2.IExecuteResult {
	res, err := c.service.EventsForDay(c.userId, c.key)
	if err != nil {
		result := &Result2.GetExecuteResult{}
		result.SetError(err)
		return result
	}

	return &Result2.GetExecuteResult{
		Events: res,
	}
}
