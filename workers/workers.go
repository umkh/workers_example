package workers

import (
	"context"
	"golang.org/x/sync/errgroup"
	"log"
	"workspace/workers_example/workers/down_worker"
)

// Worker ...
type Worker interface {
	Run(ctx context.Context) error
}

// WorkerPool ...
type WorkerPool struct {
	workers []Worker
}

// NewWorkerPool ...
func NewWorkerPool() *WorkerPool {
	wp := &WorkerPool{}
	wp.AddWorker(down_worker.NewDownWorker())
	return wp
}

// Run ...
func (wp *WorkerPool) Run(ctx context.Context) error {
	g, gCtx := errgroup.WithContext(ctx)

	for i, w := range wp.workers {
		log.Printf("Worker run %d", i)

		wr := w
		g.Go(func() error {
			return wr.Run(gCtx)
		})
	}

	return g.Wait()
}

// AddWorker ...
func (wp *WorkerPool) AddWorker(w Worker) {
	wp.workers = append(wp.workers, w)
}
