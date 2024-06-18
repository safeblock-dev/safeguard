package safeguard

import (
	"fmt"
	"os"
	"strings"
)

func Report(errs ...error) {
	const (
		colorReset = "\033[0m"
		colorRed   = "\033[31m"
		title      = colorRed + "[---------- FAIL ----------]" + colorReset
		end        = title
	)

	for _, err := range errs {
		fmt.Printf("\n%s\n\n\t%s\n\n%s\n", title, printStackTrace(err), end) //nolint: forbidigo
	}

	os.Exit(1)
}

func printStackTrace(err error) string {
	return strings.ReplaceAll(err.Error(), "\n", "\n\t")
}
