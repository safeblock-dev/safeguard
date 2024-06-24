package safeguard_test

import (
	"testing"

	"github.com/safeblock-dev/safeguard"
	"github.com/stretchr/testify/require"
)

func TestCatch(t *testing.T) {
	t.Parallel()

	const expectedPanic = "example panic"

	t.Run("no error and no panic", func(t *testing.T) {
		t.Parallel()
		var called bool
		safeguard.Catch(func() error {
			return nil
		}, func(_ error) {
			called = true
		})
		require.False(t, called, "handler should not be called if no error or panic occurs")
	})

	t.Run("with panic", func(t *testing.T) {
		t.Parallel()

		var capturedErr error
		func() {
			var err error
			defer safeguard.Catch(func() error {
				return err
			}, func(err error) {
				capturedErr = err
			})

			panic(expectedPanic)
		}()
		require.Error(t, capturedErr)
		require.Contains(t, capturedErr.Error(), expectedPanic, "handler should capture the panic message")
	})
}
