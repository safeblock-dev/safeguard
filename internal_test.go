package safeguard // nolint: testpackage

import (
	"errors"
	"testing"

	"github.com/safeblock-dev/werr"
	"github.com/stretchr/testify/require"
)

func TestFilterSkipErrors(t *testing.T) {
	t.Parallel()

	err1, err2, err3 := errors.New("1"), errors.New("2"), errors.New("3")

	t.Run("when does not contains skip error", func(t *testing.T) {
		t.Parallel()

		errs := []error{err1, err2}
		exp := errs

		require.Equal(t, exp, filterSkipErrors(errs, err3))
	})

	t.Run("when skip all errors", func(t *testing.T) {
		t.Parallel()

		errs := []error{err1}

		require.Empty(t, filterSkipErrors(errs, err1))
	})

	t.Run("when contains simple skip error", func(t *testing.T) {
		t.Parallel()

		errs := []error{err1, err2, err3}
		exp := []error{err1, err2}

		require.Equal(t, exp, filterSkipErrors(errs, err3))
	})

	t.Run("when contains join skip error", func(t *testing.T) {
		t.Parallel()

		errs := []error{err1, errors.Join(err2, err3)}
		exp := []error{err1, err2}

		require.Equal(t, exp, filterSkipErrors(errs, err3))
	})

	t.Run("when contains wrap join skip error", func(t *testing.T) {
		t.Parallel()

		errs := []error{err1, werr.Wrap(errors.Join(err2, err3))}
		exp := []error{err1, err2}

		require.Equal(t, exp, filterSkipErrors(errs, err3))
	})

	t.Run("when contains wrap join skip error", func(t *testing.T) {
		t.Parallel()

		errs := []error{err1, werr.Wrap(errors.Join(err2, err3))}
		exp := []error{err1, err2}

		require.Equal(t, exp, filterSkipErrors(errs, err3))
	})
}

func TestProcessOptions(t *testing.T) {
	t.Parallel()

	t.Run("no options", func(t *testing.T) {
		t.Parallel()

		require.NotPanics(t, func() { processOptions([]error{}) })
	})

	t.Run("with func no error", func(t *testing.T) {
		t.Parallel()

		called := false
		opt := func() { called = true }

		processOptions([]error{}, opt)
		require.True(t, called)
	})

	t.Run("with func error", func(t *testing.T) {
		t.Parallel()

		var receivedError error
		opt := func(err error) { receivedError = err }

		errs := []error{errors.New("test error")}
		processOptions(errs, opt)
		require.Equal(t, errs[0], receivedError)
	})

	t.Run("with func variadic error", func(t *testing.T) {
		t.Parallel()

		var receivedErrors []error
		opt := func(errs ...error) { receivedErrors = errs }

		errs := []error{errors.New("test error 1"), errors.New("test error 2")}
		processOptions(errs, opt)
		require.Equal(t, errs, receivedErrors)
	})

	t.Run("with func slice error", func(t *testing.T) {
		t.Parallel()

		var receivedErrors []error
		opt := func(errs []error) { receivedErrors = errs }

		errs := []error{errors.New("test error 1"), errors.New("test error 2")}
		processOptions(errs, opt)
		require.Equal(t, errs, receivedErrors)
	})

	t.Run("with skipErr", func(t *testing.T) {
		t.Parallel()

		skip := errors.New("skip this error")
		err1 := errors.New("test error 1")
		err2 := errors.New("test error 2")

		var receivedErrors []error
		filterFunc := func(errs []error) {
			receivedErrors = errs
		}

		processOptions([]error{err1, skip, err2}, SkipErr(skip), filterFunc)
		require.Equal(t, []error{err1, err2}, receivedErrors)
	})

	t.Run("with same error", func(t *testing.T) {
		t.Parallel()

		err1 := errors.New("test error 1")
		err2 := errors.New("test error 2")

		var receivedErrors []error
		filterFunc := func(errs []error) {
			receivedErrors = errs
		}

		processOptions([]error{err1}, err2, filterFunc)
		require.Equal(t, []error{err1, err2}, receivedErrors)
	})
}
