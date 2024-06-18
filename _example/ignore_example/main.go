package main

import (
	"errors"

	"github.com/safeblock-dev/safeguard"
)

var skipErr = errors.New("example error")

func main() {
	var err error
	defer safeguard.Catch(func() error {
		return err
	}, skipErr, safeguard.Report)

	err = skipErr
	panic("not skip panic")
}
