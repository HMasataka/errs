package errs

// WithMessage returns error with the given message data.
// This data is typically used for user-facing messages and can be any type,
// often a string or a map for internationalization.
//
// Example with string message:
//
//	err := New("user.not_found").WithMessage("User not found")
//
// Example with i18n data:
//
//	err := New("user.not_found").WithMessage(map[string]string{
//		"en": "User not found",
//		"ja": "ユーザーが見つかりません",
//	})
//
// The message data can be extracted using the Message() or MessageOr() functions.
func (e *Error) WithMessage(data any) *Error {
	e.message = data
	return e
}

// Message attempts to extract structured message data from an error.
//
// Example:
//
//	err := New("user.not_found").WithMessage(map[string]string{
//		"en": "User not found",
//	})
//
//	if msg, ok := Message[map[string]string](err); ok {
//		fmt.Println(msg["en"]) // "User not found"
//	}
//
// This function is particularly useful for extracting structured message data
// such as translation maps or validation error details.
func Message[T any](err error) (T, bool) {
	var zero T

	if e, ok := err.(*Error); ok && e.message != nil {
		if data, ok := e.message.(T); ok {
			return data, true
		}
	}

	return zero, false
}

// MessageOr extracts typed message data from an error with a fallback value.
//
// Example:
//
//	msg := MessageOr(err, "Unknown error")
//	// msg will be the extracted message or "Unknown error" if extraction fails
//
// This function provides a convenient way to safely extract message data
// without needing to check the boolean return value.
func MessageOr[T any](err error, fallback T) T {
	if data, ok := Message[T](err); ok {
		return data
	}

	return fallback
}
