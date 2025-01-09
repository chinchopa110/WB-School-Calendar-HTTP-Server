package Result

type IExecuteResult interface {
	SetError(Error error)
	GetError() error
}
