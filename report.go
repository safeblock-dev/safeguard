package safeguard

import (
	"fmt"
	"os"
	"strings"
)

const (
	colorReset = "\033[0m"  // ANSI escape code for resetting text color
	colorRed   = "\033[31m" // ANSI escape code for red text
	colorGreen = "\033[32m" // ANSI escape code for green text
)

// ReportAndExit prints error reports with color-coded formatting and exits the program.
// If no errors are provided, it prints a success message and exits with code 0.
// If errors are provided, it prints each error with a formatted title and exits with code 1.
// nolint: forbidigo
func ReportAndExit(errs ...error) {
	const (
		fTitle = colorRed + "[---------- FAIL[%d] ----------]" + colorReset
		fEnd   = fTitle
		sTitle = colorGreen + "\n---------- SUCCESS ----------" + colorGreen
	)
	if len(errs) == 0 {
		fmt.Println(sTitle) // Print success message
		os.Exit(0)          // Exit with success code
	}

	// Print each error with formatted title and exit with error code
	for i, err := range errs {
		fmt.Printf("\n%s\n\n\t%s\n\n%s\n",
			fmt.Sprintf(fTitle, i), printStackTrace(err), fmt.Sprintf(fEnd, i),
		)
	}

	os.Exit(1) // Exit with error code
}

// printStackTrace formats an error's stack trace for printing purposes.
func printStackTrace(err error) string {
	return strings.ReplaceAll(err.Error(), "\n", "\n\t")
}
