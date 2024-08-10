package errors

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
)

type WrapError struct {
	ErrorType ErrorType
	err       error
	trigger   error
	file      string
	line      uint
}

func New(errorId ErrorId, errType ErrorType) *WrapError {
	file, line := whereIsError()
	err := errors.New(string(errorId))
	return newWrapError(err, file, line)
}

func NewWrap(errorId ErrorId, trigger error) *WrapError {
	file, line := whereIsError()
	err := errors.New(string(errorId))
	wrapErr, ok := trigger.(WrapError)
	if !ok {
		return newWrapError(err, file, line).Wrap(trigger)
	}
	return newWrapError(err, file, line).Wrap(trigger).WithType(wrapErr.ErrorType)
}

func NewWrapWithType(trigger error, errType ErrorType, errorId ErrorId) *WrapError {
	file, line := whereIsError()
	err := errors.New(string(errorId))
	return newWrapError(err, file, line).Wrap(trigger).WithType(errType)
}

func newWrapError(err error, file string, line uint) *WrapError {
	return &WrapError{
		err:  err,
		file: file,
		line: line,
	}
}

func whereIsError() (string, uint) {
	_, file, line, _ := runtime.Caller(2)
	return file, uint(line)
}

func (e WrapError) Error() string {
	return e.format()
}

func (e *WrapError) Wrap(trigger error) *WrapError {
	e.trigger = trigger
	return e
}

func (e *WrapError) WithType(errType ErrorType) *WrapError {
	e.ErrorType = errType
	return e
}

func (e *WrapError) Unwrap() error {
	return e.trigger
}

func (e *WrapError) format() string {
	var errStr strings.Builder
	errStr.WriteString(fmt.Sprintf("ERROR: %s\n", e.err.Error()))
	errStr.WriteString(fmt.Sprint("file:", e.file))
	errStr.WriteString(" ")
	errStr.WriteString(fmt.Sprint("line:", e.line))
	if trigger := e.Unwrap(); trigger != nil {
		errStr.WriteString("\n")
		errStr.WriteString(trigger.Error())
	}
	return errStr.String()
}
