package error

import "fmt"

// OperatorError 操作错误
type OperatorError struct {
	Code    Code   `json:"code"`
	Message string `json:"message"`
}

func (e *OperatorError) Error() string {
	return e.Message
}

func NewOperatorError(code Code, message ...string) *OperatorError {
	if len(message) == 0 {
		message[0] = code.String()
	}

	return &OperatorError{Code: code, Message: message[0]}
}

func NewCreateError(model, msg string) *OperatorError {
	return &OperatorError{
		Code:    InsertError,
		Message: fmt.Sprintf("model %s create error:%s", model, msg),
	}
}
