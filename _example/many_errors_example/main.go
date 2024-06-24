package main

import (
	"errors"
	"fmt"

	"github.com/safeblock-dev/safeguard"
	"github.com/safeblock-dev/wr/taskgroup"
)

func main() {
	err := errors.New("example error 1")
	err2 := fmt.Errorf("example error 2")
	err3 := taskgroup.ErrSignal

	defer func() {
		safeguard.Catch(
			func() error { return err }, err2, err3,
			safeguard.SkipErr(taskgroup.ErrSignal),
			safeguard.ReportAndExit,
		)
	}()

	panic("example panic")
}
