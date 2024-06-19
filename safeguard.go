package safeguard

import (
	"errors"
	"fmt"
	"reflect"
	"slices"

	"github.com/safeblock-dev/werr"
)

// skipErr is a specific error type used to indicate errors that should be skipped.
type skipErr struct {
	error
}

// Catch captures and handles errors and panics that occur during the execution of the provided error.
// The second parameter 'options' allows for passing error handlers and specific errors to be ignored.
// Supported types for 'options' include:
// - func(error): a handler for single errors
// - func(...error): a handler for multiple errors
// - func([]error): a handler for a slice of errors
// - error: specific errors to be ignored.
func Catch(exception error, options ...any) {
	errs := CollectErrors(exception, werr.PanicToError(recover()))

	// Process each option
	processOptions(errs, options...)
}

// CatchFn captures and handles errors and panics that occur during the execution of the provided function.
// The second parameter 'options' allows for passing error handlers and specific errors to be ignored.
// Supported types for 'options' include:
// - func(error): a handler for single errors
// - func(...error): a handler for multiple errors
// - func([]error): a handler for a slice of errors
// - error: specific errors to be ignored.
func CatchFn(f func() error, options ...any) {
	errs := CollectErrors(f(), werr.PanicToError(recover()))

	// Process each option
	processOptions(errs, options...)
}

// CollectErrors collects the provided errors and any panic converted to an error, removing nil values.
func CollectErrors(errs ...error) []error {
	return slices.DeleteFunc(errs, func(err error) bool { return err == nil })
}

// processOptions processes the provided options to handle errors accordingly.
func processOptions(errs []error, options ...any) {
	for _, opt := range options {
		switch x := opt.(type) {
		case func(error):
			for _, err := range errs {
				x(err)
			}
		case func(...error):
			x(errs...)
		case func([]error):
			x(errs)
		case skipErr:
			errs = filterErrors(errs, x)
		case error:
			errs = append(errs, x)
		default:
			panic(fmt.Sprintf("safeguard: unsupported option type provided %v", reflect.TypeOf(opt)))
		}
	}
}

// filterErrors filters out specific errors from the error list.
func filterErrors(errs []error, specificErr skipErr) []error {
	var filtered []error
	for _, err := range errs {
		if !errors.Is(err, specificErr.error) {
			filtered = append(filtered, err)
		}
	}

	return filtered
}

// SkipErr creates a skipErr from a standard error.
func SkipErr(err error) skipErr { //nolint: revive
	return skipErr{error: err}
}
