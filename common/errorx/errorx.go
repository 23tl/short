package errorx

const defaultCode = 10001

type CodeError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type BaseCodeError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewCodeError(code int, message string) error {
	return &CodeError{Code: code, Message: message}
}

func DefaultCodeError(message string) error {
	return NewCodeError(defaultCode, message)
}

func (e *CodeError) Error() string {
	return e.Message
}

func (e *CodeError) Date() *BaseCodeError {
	return &BaseCodeError{Code: e.Code, Message: e.Message}
}
