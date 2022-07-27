package utils

// ErrorResult 错误结果
type ErrorResult struct {
	data interface{}
	err  error
}

func NewErrorResult(data interface{}, err error) *ErrorResult {
	return &ErrorResult{data: data, err: err}
}

// Unwrap 错误处理
func (r *ErrorResult) Unwrap() interface{} {
	if r.err != nil {
		panic(r.err.Error())
	}

	return r.data
}

type BindFunc func(interface{}) error

// Exec 执行
func Exec(f BindFunc, v interface{}) *ErrorResult {
	return NewErrorResult(v, f(v))
}
