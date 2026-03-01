package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup, ch <-chan string, workerNo int) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker: %d - stopped\n", workerNo)
			return
		case j, ok := <-ch:
			if !ok {
				return
			}
			time.Sleep(2 * time.Second)
			fmt.Printf("Worker: %d, %s\n", workerNo, j)
		}
	}

}

func workerPool() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	ch := make(chan string)

	var wg sync.WaitGroup
	workersCount := 3
	wg.Add(workersCount)

	for i := range workersCount {
		go worker(ctx, &wg, ch, i)
	}

	go func() {
		for i := range 10 {
			ch <- "job_" + strconv.Itoa(i)
			time.Sleep(2 * time.Second)

		}
		close(ch)
	}()

	wg.Wait()

	fmt.Printf("All %d workers stopped, shutting down...", workersCount)
}
