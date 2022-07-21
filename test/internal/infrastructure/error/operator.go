package error

import (
	"strings"
)

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
		message = append(message, code.String())
	}

	return &OperatorError{Code: code, Message: message[0]}
}

func NewCreateError(model, msg string) *OperatorError {
	builder := strings.Builder{}
	builder.WriteString("model:")
	builder.WriteString(model)
	builder.WriteString(" create error:")
	builder.WriteString(msg)

	return &OperatorError{
		Code:    InsertError,
		Message: builder.String(),
	}
}
