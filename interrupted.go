package interrupted

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// Wait waits for an interrupt signal to be received, then cancels
// the context and waits for the wait group to finish. If a second interrupt
// signal is received, the program is terminated.
func Wait(function func(context.Context, *sync.WaitGroup)) {
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	function(ctx, wg)
	go catchInterrupt(ctx, cancel)

	wg.Wait()
	os.Exit(0)
}

func catchInterrupt(ctx context.Context, cancel context.CancelFunc) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Wait for interrupt signal or context cancellation.
	select {
	case <-interrupt:
		cancel()
	case <-ctx.Done():
	}

	// Wait for a second interrupt signal to force an exit.
	<-interrupt
	os.Exit(2)
}
