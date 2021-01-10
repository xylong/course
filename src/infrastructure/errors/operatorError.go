package errors

import "fmt"

const (
	CreateError = 20001 // 实体新增失败
	UpdateError = 20002 // 实体更新失败
)

type OperatorError struct {
	Code int
	Msg  string
}

func (o *OperatorError) Error() string {
	return o.Msg
}

func NewOperator(code int, msg string) *OperatorError {
	return &OperatorError{Code: code, Msg: msg}
}

func NewCreateError(model string, err string) *OperatorError {
	return &OperatorError{
		Code: CreateError,
		Msg:  fmt.Sprintf("model %s create error:%s", model, err),
	}
}

func NewUpdateError(model string, err string) *OperatorError {
	return &OperatorError{
		Code: CreateError,
		Msg:  fmt.Sprintf("model %s update error:%s", model, err),
	}
}
