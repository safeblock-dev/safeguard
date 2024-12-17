package safeguard

import (
	"log"
	"os"
	"strings"
)

var StdLogger = log.New(os.Stderr, "", log.LstdFlags) //nolint: gochecknoglobals // std logger.

// Report prints error reports with color-coded formatting.
func Report(errs ...error) {
	if len(errs) == 0 {
		//nolint: errcheck // don't need.
		StdLogger.Writer().Write([]byte("\n"))
		StdLogger.Println("no errors")
	}

	// Print each error with formatted title and exit with error code
	for _, err := range errs {
		//nolint: errcheck // don't need.
		StdLogger.Writer().Write([]byte("\n"))
		StdLogger.Println(strings.ReplaceAll(err.Error(), "\n", "\n\t"))
	}
}

// ReportAndExit prints error reports with color-coded formatting and exits the program.
// If no errors are provided, it prints a success message and exits with code 0.
// If errors are provided, it prints each error with a formatted title and exits with code 1.
// nolint: forbidigo
func ReportAndExit(errs ...error) {
	Report(errs...)
	Exit(errs...) // Exit with error code
}
