package util

type ErrorResult struct {
	data interface{}
	err  error
}

func NewErrorResult(data interface{}, err error) *ErrorResult {
	return &ErrorResult{data: data, err: err}
}

func (r *ErrorResult) Unwrap() interface{} {
	if r.err != nil {
		panic(r.err.Error())
	}
	return r.data
}

type BindFunc func(any interface{}) error

func Exec(bindFunc BindFunc, any interface{}) *ErrorResult {
	err := bindFunc(any)
	return NewErrorResult(any, err)
}
