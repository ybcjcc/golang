package xcode

import (
	"net/http"
)

type Code struct {
	httpStatus int
	code       int
	message    string
}

func (c Code) Code() int {
	return c.code
}

func (c Code) HttpStatus() int {
	return c.httpStatus
}

func (c Code) String() string {
	return c.message
}

func New(code int) XCode {
	// 如果是已定义的 code，直接返回；否则构建一个
	for _, c := range allCode {
		if c.code == code {
			return c
		}
	}

	return NewWithMessage(code, InternalServerError.message)
}

func NewWithMessage(code int, message string) XCode {
	// 检查错误码
	if MinReservedCode <= code && code <= MaxReservedCode {
		return ErrorCodeFailed
	}

	return &Code{
		httpStatus: http.StatusInternalServerError,
		code:       code,
		message:    message,
	}
}

func WithMessage(xcode XCode, message string) XCode {
	return &Code{
		httpStatus: xcode.HttpStatus(),
		code:       xcode.Code(),
		message:    message,
	}
}

func NewUseMessage(message string) XCode {
	c := InternalServerError
	c.message = message

	return c
}
