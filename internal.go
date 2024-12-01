package safeguard

import (
	"errors"
	"fmt"
	"reflect"
)

// skipError is a specific error type used to indicate errors that should be skipped.
type skipError struct {
	error
}

// FilterSkipErrors filters out specific errors from the error list, including joined errors.
func filterSkipErrors(errs []error, skip error) []error {
	filtered := make([]error, 0, len(errs))
	for _, err := range errs {
		// Check if the error is a joined error
		var joinedErr interface{ Unwrap() []error }
		if errors.As(err, &joinedErr) {
			subErrors := joinedErr.Unwrap()
			for _, subErr := range subErrors {
				if !errors.Is(subErr, skip) {
					filtered = append(filtered, subErr)
				}
			}
		} else if !errors.Is(err, skip) {
			filtered = append(filtered, err)
		}
	}

	return filtered
}

// processOptions processes the provided options to handle errors accordingly.
func processOptions(errs []error, options ...any) {
	for _, opt := range options {
		switch x := opt.(type) {
		case func():
			x()
		case func(error):
			for _, err := range errs {
				x(err)
			}
		case func(...error):
			x(errs...)
		case func([]error):
			x(errs)
		case skipError:
			errs = filterSkipErrors(errs, x.error)
		case error:
			errs = append(errs, x)
		default:
			panic(fmt.Sprintf("safeguard: unsupported option type provided %v", reflect.TypeOf(opt)))
		}
	}
}
