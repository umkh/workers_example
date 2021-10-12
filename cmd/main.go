package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"workspace/workers_example/workers"
)

func main() {
	ctx, shutdown := context.WithCancel(context.Background())

	wp := workers.NewWorkerPool()
	go func(ctx context.Context, wp *workers.WorkerPool) {
		if wpErr := wp.Run(ctx); wpErr != nil {
			log.Printf("Worker pool error %s", wpErr.Error())
		}
	}(ctx, wp)

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	shutdown()
	fmt.Println("Graceful Shutdown ...")
}
