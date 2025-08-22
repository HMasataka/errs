package errs

import (
	"encoding/json"

	"github.com/HMasataka/stalker"
)

// Trace provides a way to capture and represent the stack trace of an error.
func (e *Error) Trace() string {
	b, _ := json.Marshal(e.frame)
	return string(b)
}

// WithCallerStack captures the current stack frame and associates it with the error.
func (e *Error) WithCallerStack() *Error {
	e.frame = stalker.Wrap(e.frame, stalker.SkipFrame(3))
	return e
}
