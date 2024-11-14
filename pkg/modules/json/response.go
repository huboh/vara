package json

import (
	"runtime/debug"
)

// Status is the response status.
type Status string

// recognized response Status
const (
	StatusError   = Status("error")
	StatusSuccess = Status("success")
)

// Response represent the response object.
type Response struct {
	// Status is the response status. either "success" or "error".
	Status Status `json:"status"`

	// Message is the response message. defaults to http.StatusText(Code).
	Message string `json:"message"`

	// StatusCode is the http response status code.
	StatusCode int `json:"statusCode"`

	// Data is the response data.
	Data any `json:"data,omitempty"`

	// Error is the response error. this field is omitted from the response if it is nil.
	Error *Error `json:"error,omitempty"`
}

// Error represent the response error object.
type Error struct {
	// Name is the name of the error.
	Name string `json:"name,omitempty"`

	// Cause is the error's cause.
	Cause string `json:"cause,omitempty"`

	// Stack is the error stack trace.
	Stack string `json:"stack,omitempty"`

	// Message is the error message
	Message string `json:"message,omitempty"`
}

func NewError(name string, msg string, cause string, stack string) *Error {
	return &Error{
		Message: msg,
		Cause:   cause,
		Stack:   stack,
		Name:    name,
	}
}

// ErrorFromErr returns a new Error only if the passed in err is not nil. it returns nil otherwise
func ErrorFromErr(err error, name string, cause string) *Error {
	if err == nil {
		return nil
	}

	return NewError(name, err.Error(), cause, string(debug.Stack()))
}

// Error makes ResponseError meets the error interface
func (re *Error) Error() string {
	return re.Message
}
