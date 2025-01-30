package Commands

import (
	"WB2/internal/Presentation/RestAPI/Commands/Result"
)

type ICommand interface {
	Execute() Result.IExecuteResult
}
