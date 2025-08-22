package errs

// WithHTTPStatusCode returns the error with the specified HTTP status code.
// This is useful for web applications that need to map errors to appropriate
// HTTP response codes.
//
// Example:
//
//	err := New("user.not_found").
//		WithHTTPStatusCode(404).
//		WithMessage("User not found")
func (e *Error) WithHTTPStatusCode(statusCode int) *Error {
	e.status = statusCode
	return e
}

// HTTPStatusCode returns the HTTP status code associated with this error.
// Returns 0 if no HTTP status code was set.
//
// This method is typically used by web frameworks or middleware to
// determine the appropriate HTTP response code for an error.
func (e *Error) HTTPStatusCode() int {
	return e.status
}

// HTTPStatus extracts the HTTP status code from any error.
// If the error is an Error with a status code, returns that code.
// Otherwise returns 0.
//
// This function enables HTTP status code extraction from any error in
// an error chain, making it useful for middleware and error handlers.
//
// Example:
//
//	status := errorsx.HTTPStatus(err)
//	if status != 0 {
//		w.WriteHeader(status)
//	} else {
//		w.WriteHeader(500) // Default to internal server error
//	}
//
// Returns 0 if no HTTP status is found or if err is nil.
func HTTPStatus(err error) int {
	if e, ok := err.(*Error); ok && e.status != 0 {
		return e.status
	}
	return 0
}
