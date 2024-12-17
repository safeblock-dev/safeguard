package safeguard_test

import (
	"bytes"
	"errors"
	"os"
	"testing"

	"github.com/safeblock-dev/safeguard"
	"github.com/stretchr/testify/require"
)

// nolint: paralleltest
func TestReport(t *testing.T) {
	t.Run("no errors", func(t *testing.T) {
		buf := new(bytes.Buffer)
		safeguard.StdLogger.SetOutput(buf)
		safeguard.Report()
		safeguard.StdLogger.SetOutput(os.Stderr)

		require.Contains(t, buf.String(), "no errors")
	})

	t.Run("with errors", func(t *testing.T) {
		errs := []error{
			errors.New("test error 1"),
			errors.New("test error 2"),
		}

		buf := new(bytes.Buffer)
		safeguard.StdLogger.SetOutput(buf)
		safeguard.Report(errs...)
		safeguard.StdLogger.SetOutput(os.Stderr)

		for _, err := range errs {
			require.Contains(t, buf.String(), err.Error())
		}
	})
}
