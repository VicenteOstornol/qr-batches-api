package errors

import (
	"fmt"
	"runtime"

	"github.com/ansel1/merry"
)

func New(message string) error {
	return merry.New(message)
}

func Wrap(err error, message string) error {
	return merry.Prepend(err, message)
}

func WithHTTPCode(err error, status int) merry.Error {
	return merry.WithHTTPCode(err, status)
}

func Wrapf(err error, format string, args ...interface{}) error {
	return merry.Prependf(err, format, args...)
}

func Is(err error, origingals ...error) bool {
	return merry.Is(err, origingals...)
}

type ErrorCause struct {
	errMsg string
	Values map[string]interface{}
	Trace  []Frame
}

type Frame struct {
	Function string `json:"function,omitempty"`
	File     string `json:"file,omitempty"`
	Line     int    `json:"line,omitempty"`
}

func (ec *ErrorCause) Error() string {
	return ec.errMsg
}

func GetCauseFromError(err error) *ErrorCause {
	return &ErrorCause{
		errMsg: err.Error(),
		Values: makeValues(err),
		Trace:  makeTrace(err),
	}
}

func makeValues(err error) map[string]interface{} {
	values := make(map[string]interface{})
	for k, v := range merry.Values(err) {
		values[fmt.Sprintf("%v", k)] = v
	}
	return values
}

func makeTrace(err error) []Frame {
	ff := []Frame{}
	stackPointer := merry.Stack(err)
	frames := runtime.CallersFrames(stackPointer)
	for {
		frame, more := frames.Next()
		if !more {
			break
		}
		ff = append(ff, Frame{
			Function: frame.Function,
			File:     frame.File,
			Line:     frame.Line,
		})
	}
	return ff
}
