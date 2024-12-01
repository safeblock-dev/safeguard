package safeguard

import (
	"slices"

	"github.com/safeblock-dev/werr"
)

// Catch captures and handles errors and panics that occur during the execution of the provided function.
// The second parameter 'options' allows for passing error handlers and specific errors to be ignored.
// Supported types for 'options' include:
// - func(error): a handler for single errors
// - func(...error): a handler for multiple errors
// - func([]error): a handler for a slice of errors
// - error: specific errors to be ignored.
func Catch(fn func() error, options ...any) {
	errs := CollectErrors(fn(), werr.PanicToError(recover()))

	// Process each option
	processOptions(errs, options...)
}

// CollectErrors collects the provided errors and any panic converted to an error, removing nil values.
func CollectErrors(errs ...error) []error {
	return slices.DeleteFunc(errs, func(err error) bool { return err == nil })
}

// SkipErr creates a skipError from a standard error.
func SkipErr(err error) skipError { //nolint: revive
	return skipError{error: err}
}
