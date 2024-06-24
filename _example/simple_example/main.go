package main

import (
	"context"
	"errors"
	"log"
	"os"
	"syscall"
	"time"

	"github.com/safeblock-dev/safeguard"
	"github.com/safeblock-dev/wr/taskgroup"
)

func main() {
	ctx := context.Background()

	var err error
	defer func() {
		safeguard.Catch(
			func() error { return err },
			safeguard.SkipErr(taskgroup.ErrSignal),
			safeguard.ReportAndExit,
		)
	}()

	tasks := taskgroup.New()
	tasks.Add(taskgroup.SignalHandler(ctx, os.Interrupt, syscall.SIGINT, syscall.SIGTERM))
	tasks.AddContext(func(ctx context.Context) error {
		select {
		case <-ctx.Done():
			return nil
		case <-time.NewTimer(5 * time.Second).C:
			return errors.New("example fail")
		}
	}, taskgroup.SkipInterruptCtx())

	log.Println("press CTR + C")
	err = tasks.Run()
}
