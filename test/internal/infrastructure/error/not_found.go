package error

// NotFoundError 未找到错误
type NotFoundError struct {
	Code    Code   `json:"code"`
	Message string `json:"message"`
}

func (e *NotFoundError) Error() string {
	return e.Message
}

func NewNotFoundError(code Code, message ...string) *NotFoundError {
	if len(message) == 0 {
		message[0] = code.String()
	}

	return &NotFoundError{Code: code, Message: message[0]}
}

func NewNoIDError(entity string) *NotFoundError {
	return &NotFoundError{
		Code:    ZeroIEntityD,
		Message: entity + ":" + ZeroIEntityD.String(),
	}
}

func NewNoDataError(entity string) *NotFoundError {
	return &NotFoundError{
		Code:    NotFoundData,
		Message: entity + ":" + NotFoundData.String(),
	}
}
