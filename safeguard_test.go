package safeguard_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/safeblock-dev/safeguard"
	"github.com/safeblock-dev/werr"
	"github.com/stretchr/testify/require"
)

func TestCatch(t *testing.T) {
	t.Parallel()

	const expectedPanic = "example panic"
	expectedErr := errors.New("example error")

	t.Run("no error and no panic", func(t *testing.T) {
		t.Parallel()
		var called bool
		safeguard.Catch(nil, func(_ error) {
			called = true
		})
		require.False(t, called, "handler should not be called if no error or panic occurs")
	})

	t.Run("with error", func(t *testing.T) {
		t.Parallel()

		var capturedErr error
		safeguard.Catch(expectedErr, func(err error) {
			capturedErr = err
		})
		require.Equal(t, expectedErr, capturedErr, "handler should capture the provided error")
	})

	t.Run("with panic", func(t *testing.T) {
		t.Parallel()

		var capturedErr error
		func() {
			defer safeguard.Catch(nil, func(err error) {
				capturedErr = err
			})

			panic(expectedPanic)
		}()
		require.Error(t, capturedErr)
		require.Contains(t, capturedErr.Error(), expectedPanic, "handler should capture the panic message")
	})

	t.Run("multiple options", func(t *testing.T) {
		t.Parallel()

		var capturedErr1, capturedErr2 error
		safeguard.Catch(expectedErr, func(err error) {
			capturedErr1 = err
		}, func(err error) {
			capturedErr2 = err
		})
		require.Equal(t, expectedErr, capturedErr1, "first handler should capture the error")
		require.Equal(t, expectedErr, capturedErr2, "second handler should capture the error")
	})

	t.Run("with panic and error", func(t *testing.T) {
		t.Parallel()

		var capturedErrs []error
		func() {
			defer safeguard.Catch(expectedErr, func(err error) {
				capturedErrs = append(capturedErrs, err)
			})

			panic(expectedPanic)
		}()
		require.Len(t, capturedErrs, 2)
		require.Equal(t, expectedErr, capturedErrs[0])
		require.Contains(t, capturedErrs[1].Error(), expectedPanic)
	})

	t.Run("second error", func(t *testing.T) {
		t.Parallel()

		expected2Err := errors.New("example error 2")
		expected3Err := errors.New("example error 3")
		var capturedErrs []error
		safeguard.Catch(expectedErr, expected2Err, expected3Err, func(err error) {
			capturedErrs = append(capturedErrs, err)
		})
		require.Len(t, capturedErrs, 3)
		require.Equal(t, expectedErr, capturedErrs[0])
		require.Equal(t, expected2Err, capturedErrs[1])
		require.Equal(t, expected3Err, capturedErrs[2])
	})

	t.Run("handle slice error handler", func(t *testing.T) {
		t.Parallel()
		expectedErrs := []error{
			errors.New("error 1"),
			errors.New("error 2"),
		}

		var called bool
		func() {
			defer safeguard.Catch(expectedErrs[0], func(errs []error) {
				called = true
				require.Equal(t, expectedErrs[0], werr.Cause(errs[0]))
				require.Equal(t, expectedErrs[1], werr.Cause(errs[1]))
			})

			panic(expectedErrs[1])
		}()

		require.True(t, called)
	})

	t.Run("when simple func option", func(t *testing.T) {
		t.Parallel()

		var called bool
		safeguard.Catch(expectedErr, func() {
			called = true
		})
		require.True(t, called)
	})

	t.Run("unsupported option type", func(t *testing.T) {
		t.Parallel()

		require.Panics(t, func() {
			safeguard.Catch(nil, 123) // Passing unsupported option type (int)
		}, "safeguard: unsupported option type provided")
	})

	t.Run("when skip error", func(t *testing.T) {
		t.Parallel()

		expected2Err := errors.New("example error 2")
		skippedErr := fmt.Errorf("example error 3") //nolint: perfsprint
		var capturedErrs []error
		safeguard.Catch(expectedErr, expected2Err, skippedErr,
			safeguard.SkipErr(skippedErr),
			func(err error) { capturedErrs = append(capturedErrs, err) },
		)
		require.Len(t, capturedErrs, 2)
		require.Equal(t, expectedErr, capturedErrs[0])
		require.Equal(t, expected2Err, capturedErrs[1])
	})
}

func TestCatchFn(t *testing.T) {
	t.Parallel()

	const expectedPanic = "example panic"

	t.Run("no error and no panic", func(t *testing.T) {
		t.Parallel()
		var called bool
		safeguard.CatchFn(func() error {
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
			defer safeguard.CatchFn(func() error {
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
