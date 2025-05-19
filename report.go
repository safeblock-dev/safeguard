package safeguard

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var StdLogger = log.New(os.Stderr, "", 0) //nolint: gochecknoglobals // std logger.

const reportMessagePrefix = "process finished with code "

type report struct {
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
	Status  string   `json:"status"`
	Time    string   `json:"time"`
}

// Report prints error reports with color-coded formatting.
func Report(errs ...error) {
	r := report{
		Message: reportMessagePrefix,
		Errors:  make([]string, 0),
		Status:  "unknow",
		Time:    time.Now().Format(time.RFC3339),
	}

	for _, err := range errs {
		//nolint: errcheck // don't need.
		sig, ok := err.(os.Signal)
		if ok {
			s := sig.String()
			var code int
			if code, ok = signals[s]; ok {
				r.Message += strconv.Itoa(code)
			}

			if strings.HasPrefix(s, signalPrefix) {
				r.Message += strings.TrimPrefix(s, signalPrefix)
			}

			r.Status = sig.String()
		} else {
			r.Errors = append(r.Errors, err.Error())
		}
	}

	if r.Message == reportMessagePrefix {
		if len(errs) == 0 {
			r.Message += "0"
			r.Status = "successful"
		} else {
			r.Message += "1"
			r.Status = "failed"
		}
	}

	data, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}

	StdLogger.Println(string(data))
}

// ReportAndExit prints error reports with color-coded formatting and exits the program.
// If no errors are provided, it prints a success message and exits with code 0.
// If errors are provided, it prints each error with a formatted title and exits with code 1.
// nolint: forbidigo
func ReportAndExit(errs ...error) {
	Report(errs...)
	Exit(errs...) // Exit with error code
}
