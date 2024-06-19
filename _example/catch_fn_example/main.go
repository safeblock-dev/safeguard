package main

import (
	"errors"

	"github.com/safeblock-dev/safeguard"
	"github.com/safeblock-dev/werr"
)

func main() {
	var err error
	defer safeguard.CatchFn(func() error {
		return err
	}, safeguard.ReportAndExit)

	err = errors.New("example error")
	err = werr.Wrapf(err, "wrap error 1")
	err = werr.Wrapf(err, "wrap error 2")

	panic("example panic")
}
