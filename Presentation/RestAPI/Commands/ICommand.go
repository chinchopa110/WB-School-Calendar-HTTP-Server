package Commands

import "WB2/Presentation/RestAPI/Commands/Result"

type ICommand interface {
	Execute() Result.IExecuteResult
}
