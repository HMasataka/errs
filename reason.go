package errs

import "fmt"

// WithReason returns error with a technical message.
// This message is intended for logging and debugging purposes and supports
// format string parameters similar to fmt.Sprintf.
//
// Example:
//
//	err := New("database.query_failed").
//		WithReason("Failed to execute query: %s", query)
func (e *Error) WithReason(reason string, params ...any) *Error {
	e.reason = fmt.Sprintf(reason, params...)
	return e
}
