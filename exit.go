package safeguard

import (
	"os"
)

// Exit exits the program.
// If no errors are provided, it prints a success message and exits with code 0.
// If errors are provided, it prints each error with a formatted title and exits with code 1.
func Exit(errs ...error) {
	if len(errs) > 0 {
		os.Exit(1)
	}

	os.Exit(0)
}
