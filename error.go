package bingo

// ErrorResult 错误结果
type ErrorResult struct {
	data interface{}
	err  error
}

func NewErrorResult(data interface{}, err error) *ErrorResult {
	return &ErrorResult{data: data, err: err}
}

// Unwrap 错误处理
// 如果执行有错误则抛出错误，否则返回执行结果
func (r *ErrorResult) Unwrap() interface{} {
	if r.err != nil {
		panic(message(r.err))
	}

	return r.data
}
