package down_worker

import (
	"context"
	"fmt"
	"time"
	"workspace/workers_example/pkg/utils"
)

// DownWorker ...
type DownWorker struct {
}

// NewDownWorker ...
func NewDownWorker() *DownWorker {
	return &DownWorker{}
}

// Run ...
func (w *DownWorker) Run(ctx context.Context) error {
	tr := time.NewTicker(time.Second * 5)
	var count int
	for {
		select {
		case <-tr.C:
			ctx1, cancel := context.WithCancel(ctx)
			w.Counter(count, cancel, 10)
			<-ctx1.Done()
			fmt.Println(count)
		case <-ctx.Done():
			return fmt.Errorf("context Done %w", ctx.Err())
		}
	}
}

// Counter ...
func (w *DownWorker) Counter(count int, cancel context.CancelFunc, n int) {
	for i := 0; i < n; i++ {
		go func(cancel context.CancelFunc, count int) {
			count++
			fmt.Println(utils.RandomString(8))
			defer cancel()
		}(cancel, count)
		time.Sleep(time.Second * 1)
	}
}
