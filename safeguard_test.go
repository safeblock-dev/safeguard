package safeguard_test

import (
	"errors"
	"testing"

	"github.com/safeblock-dev/safeguard"
	"github.com/safeblock-dev/werr"
	"github.com/stretchr/testify/require"
)

func TestCatch(t *testing.T) {
	t.Parallel()

	t.Run("NoErrorNoPanic", func(t *testing.T) {
		t.Parallel()
		var called bool
		safeguard.Catch(func() error {
			return nil
		}, func(_ error) {
			called = true
		})
		require.False(t, called, "handler should not be called if no error or panic occurs")
	})

	t.Run("WithError", func(t *testing.T) {
		t.Parallel()
		expectedErr := errors.New("example error")
		var capturedErr error
		safeguard.Catch(func() error {
			return expectedErr
		}, func(err error) {
			capturedErr = err
		})
		require.Equal(t, expectedErr, capturedErr, "handler should capture the returned error")
	})

	t.Run("WithPanic", func(t *testing.T) {
		t.Parallel()
		expectedPanic := "example panic"
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

	t.Run("MultipleHandlers", func(t *testing.T) {
		t.Parallel()
		expectedErr := errors.New("example error")
		var capturedErr1, capturedErr2 error
		safeguard.Catch(func() error {
			return expectedErr
		}, func(err error) {
			capturedErr1 = err
		}, func(err error) {
			capturedErr2 = err
		})
		require.Equal(t, expectedErr, capturedErr1, "first handler should capture the error")
		require.Equal(t, expectedErr, capturedErr2, "second handler should capture the error")
	})

	t.Run("WithPanicAndError", func(t *testing.T) {
		t.Parallel()
		expectedErr := errors.New("example error")
		expectedPanic := "example panic"
		var capturedErrs []error
		safeguard.Catch(func() error {
			return expectedErr
		}, func(...error) {
			capturedErrs = append(capturedErrs, expectedErr)
		}, func(...error) {
			capturedErrs = append(capturedErrs, werr.PanicToError(expectedPanic))
		})
		require.Len(t, capturedErrs, 2)
		require.Equal(t, expectedErr, capturedErrs[0], "handler should capture the returned error")
		require.Contains(t, capturedErrs[1].Error(), expectedPanic, "handler should capture the panic message")
	})

	t.Run("SkipSpecificError", func(t *testing.T) {
		t.Parallel()
		specificErr := errors.New("skip error")
		otherErr := errors.New("other error")
		var capturedErrs []error
		safeguard.Catch(func() error {
			return otherErr
		}, specificErr, func(...error) {
			capturedErrs = append(capturedErrs, otherErr)
		})
		require.Len(t, capturedErrs, 1)
		require.Equal(t, otherErr, capturedErrs[0], "handler should capture the other error")
	})

	t.Run("HandleSliceErrorHandler", func(t *testing.T) {
		t.Parallel()
		expectedErrs := []error{
			errors.New("error 1"),
			errors.New("error 2"),
		}

		var called bool
		func() {
			var err error
			defer safeguard.Catch(func() error {
				return err
			}, func(errs []error) {
				called = true
				require.Equal(t, expectedErrs[0], werr.Cause(errs[0]))
				require.Equal(t, expectedErrs[1], werr.Cause(errs[1]))
			})

			err = expectedErrs[0]
			panic(expectedErrs[1])
		}()

		require.True(t, called)
	})

	t.Run("UnsupportedOptionType", func(t *testing.T) {
		t.Parallel()

		require.Panics(t, func() {
			safeguard.Catch(func() error {
				return nil
			}, 123) // Passing unsupported option type (int)
		}, "safeguard: unsupported option type provided")
	})
}
