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
	tr := time.NewTicker(time.Second * 3)

	for {
		select {
		case <-tr.C:
			fmt.Println(utils.RandomString(8))
		case <-ctx.Done():
			return fmt.Errorf("context Done %w", ctx.Err())
		}
	}
}
