package safeguard_test

import (
	"bytes"
	"errors"
	"os"
	"testing"

	"github.com/safeblock-dev/safeguard"
	"github.com/stretchr/testify/require"
)

type signalError struct {
	sig os.Signal
}

func (err signalError) Error() string {
	return err.sig.String()
}

func (err signalError) String() string {
	return err.sig.String()
}

func (err signalError) Signal() {}

// nolint: paralleltest
func TestReport(t *testing.T) {
	t.Run("when no errors", func(t *testing.T) {
		buf := new(bytes.Buffer)
		safeguard.StdLogger.SetOutput(buf)
		safeguard.Report()
		safeguard.StdLogger.SetOutput(os.Stderr)

		require.Contains(t, buf.String(), "{\"message\":\"Process finished with code 0\",\"errors\":[],\"status\":\"successful\"", buf.String())
	})

	t.Run("when contains signal error", func(t *testing.T) {
		buf := new(bytes.Buffer)
		safeguard.StdLogger.SetOutput(buf)
		safeguard.Report(signalError{os.Interrupt})
		safeguard.StdLogger.SetOutput(os.Stderr)

		require.Contains(t, buf.String(), "{\"message\":\"Process finished with code 2\",\"errors\":[],\"status\":\"interrupt\"", buf.String())
	})

	t.Run("when contain errors", func(t *testing.T) {
		errs := []error{
			errors.New("test error 1"),
			errors.New("test error 2"),
		}

		buf := new(bytes.Buffer)
		safeguard.StdLogger.SetOutput(buf)
		safeguard.Report(errs...)
		safeguard.StdLogger.SetOutput(os.Stderr)

		require.Contains(t, buf.String(), "{\"message\":\"Process finished with code 1\",\"errors\":[\"test error 1\",\"test error 2\"],\"status\":\"failed\"", buf.String())
	})
}
