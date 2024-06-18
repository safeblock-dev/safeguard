package safeguard

import (
	"errors"

	"github.com/safeblock-dev/werr"
)

// Catch captures and handles errors and panics that occur during the execution of the provided function.
// The second parameter 'options' allows for passing error handlers and specific errors to be ignored.
// Supported types for 'options' include:
// - func(error): a handler for single errors
// - func(...error): a handler for multiple errors
// - func([]error): a handler for a slice of errors
// - func(...interface{}): a handler for variadic any type (errors in this case)
// - error: specific errors to be ignored.
func Catch(f func() error, options ...any) {
	var errs []error

	// Execute the provided function and capture its returned error
	if err := f(); err != nil {
		errs = append(errs, err)
	}

	// Recover from panic and convert it to an error
	if err := werr.PanicToError(recover()); err != nil {
		errs = append(errs, err)
	}

	// Process each option
	for _, opt := range options {
		switch h := opt.(type) {
		case func(error):
			for _, err := range errs {
				h(err)
			}
		case func(...error):
			h(errs...)
		case func([]error):
			h(errs)
		case error:
			errs = filterErrors(errs, h)
		default:
			panic("safeguard: unsupported option type provided")
		}
	}
}

// Filter out errors specified to be ignored.
func filterErrors(errs []error, filterErr error) []error {
	var filtered []error
	for _, err := range errs {
		if !errors.Is(err, filterErr) {
			filtered = append(filtered, err)
		}
	}

	return filtered
}
