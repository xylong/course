package errors

import "fmt"

const (
	NotFoundRootID    = 10001 // 没有根实体的ID
	NotFoundModelData = 10002 // 实体数据未找到
)

type NotFoundError struct {
	Code int
	Msg  string
}

func (e *NotFoundError) Error() string {
	return e.Msg
}

func NewNotFoundError(code int, msg string) *NotFoundError {
	return &NotFoundError{Code: code, Msg: msg}
}

func NewNotFoundIDError(model string) *NotFoundError {
	return &NotFoundError{Code: NotFoundRootID, Msg: fmt.Sprintf("model %s id not found", model)}

}

func NewNotFoundDataError(model string, err string) *NotFoundError {
	return &NotFoundError{Code: NotFoundModelData, Msg: fmt.Sprintf("model %s data not found %s", model, err)}
}
