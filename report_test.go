package safeguard_test

import (
	"bytes"
	"errors"
	"os"
	"testing"

	"github.com/safeblock-dev/safeguard"
	"github.com/stretchr/testify/require"
)

// nolint: paralleltest, tparallel
func TestReport(t *testing.T) {
	t.Parallel()

	t.Run("no errors", func(t *testing.T) {
		var buf bytes.Buffer
		old := os.Stdout
		defer func() { os.Stdout = old }()
		r, w, _ := os.Pipe()
		os.Stdout = w

		safeguard.Report()

		require.NoError(t, w.Close())
		_, _ = buf.ReadFrom(r)
		require.Contains(t, buf.String(), "SUCCESS")
	})

	t.Run("with errors", func(t *testing.T) {
		var buf bytes.Buffer
		old := os.Stdout
		defer func() { os.Stdout = old }()
		r, w, _ := os.Pipe()
		os.Stdout = w

		safeguard.Report(errors.New("test error 1"), errors.New("test error 2"))

		require.NoError(t, w.Close())
		_, _ = buf.ReadFrom(r)
		require.Contains(t, buf.String(), "FAIL")
	})
}
