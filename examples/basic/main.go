package main

import "github.com/HMasataka/errs"

func main() {
	err := errs.New("user.not.found").
		WithReason("User with ID %d not found", 123).
		WithHTTPStatusCode(404)

	err.WithCallerStack()

	println(err.Error(), err.Trace())
}
