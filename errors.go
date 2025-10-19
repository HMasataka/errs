package errs

import (
	"errors"

	"github.com/HMasataka/stalker"
)

type Error struct {
	id      string
	reason  string
	message any
	cause   error
	status  int
	frame   *stalker.Frame
}

// New creates a new Error with the given id and options.
// The id serves as a unique identifier for the error type and is used
// for error comparison with errors.Is().
//
// Example:
//
//	err := New("user.not_found")
//
// The id should follow a hierarchical naming convention (e.g., "domain.operation.reason")
// to facilitate error categorization and handling.
func New(id string) *Error {
	frame := stalker.NewFrame(stalker.SkipFrame(3))

	e := &Error{
		id:     id,
		reason: id, // Default reason is the ID itself
		frame:  frame,
	}

	return e
}

// Wrap creates a new Error that wraps an existing error with the specified ID.
func Wrap(err error, id string) *Error {
	if err == nil {
		return nil
	}

	// If the error is already an Error, just update the ID
	if e, ok := err.(*Error); ok {
		e.id = id
		e.frame = stalker.Wrap(e.frame, stalker.SkipFrame(3))
		return e
	}

	// Otherwise, create a new Error wrapping the original error
	frame := stalker.NewFrame(stalker.SkipFrame(3))
	return &Error{
		id:     id,
		reason: id, // Default reason is the ID itself
		cause:  err,
		frame:  frame,
	}
}

// ID returns the unique identifier of the error.
// This ID is used for error comparison and categorization.
func (e *Error) ID() string {
	return e.id
}

// Error implements the standard Go error interface.
// It returns the technical message set by WithReason(), or the error ID
// if no specific message was provided.
func (e *Error) Error() string {
	return e.reason
}

// Unwrap returns the underlying cause error, enabling Go's error unwrapping
// functionality. This allows errors.Is() and errors.As() to traverse the
// error chain to find specific error types or values.
//
// Returns nil if no underlying cause was set.
func (e *Error) Unwrap() error {
	return e.cause
}

// WithCause sets the underlying cause error for this Error.
// This allows adding a cause to an error after it has been created.
// Returns the same Error instance to enable method chaining.
//
// Example:
//
//	err := New("user.operation_failed")
//	err.WithCause(dbError)
//
// Or with method chaining:
//
//	err := New("user.operation_failed").WithCause(dbError).WithMessage("Failed to save user")
func (e *Error) WithCause(cause error) *Error {
	if e != nil {
		e.cause = cause
	}
	return e
}

// Is implements custom error comparison for errors.Is().
//
//	if errors.Is(err, New("user.not_found")) {
//		// Handle user not found error
//	}
func (e *Error) Is(target error) bool {
	if e == nil {
		return false
	}
	if e == target {
		return true
	}
	t, ok := target.(*Error)
	if !ok {
		return errors.Is(e.cause, target)
	}

	return e.id == t.id
}
