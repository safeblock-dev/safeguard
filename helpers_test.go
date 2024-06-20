package safeguard // nolint: testpackage

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFilterSkipErrors(t *testing.T) {
	t.Parallel()

	skip := fmt.Errorf("does not contains") // nolint: perfsprint
	err1, err2, err3 := errors.New("1"), errors.New("2"), errors.New("3")

	t.Run("when does not contains skip error", func(t *testing.T) {
		t.Parallel()

		errs := []error{err1, err2}
		exp := errs

		require.Equal(t, exp, filterSkipErrors(errs, SkipErr(skip)))
	})

	t.Run("when contains simple skip error", func(t *testing.T) {
		t.Parallel()

		errs := []error{err1, skip, err2}
		exp := []error{err1, err2}

		require.Equal(t, exp, filterSkipErrors(errs, SkipErr(skip)))
	})

	t.Run("when contains join skip error", func(t *testing.T) {
		t.Parallel()

		errs := []error{err1, errors.Join(skip, err2), err3}
		exp := []error{err1, err2, err3}

		require.Equal(t, exp, filterSkipErrors(errs, skip))
	})
}
