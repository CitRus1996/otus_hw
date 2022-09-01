package hw05parallelexecution

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	var wg sync.WaitGroup
	errCount := new(int32)
	taskChan := make(chan Task, len(tasks))
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i := range tasks {
		taskChan <- tasks[i]
	}
	close(taskChan)

	for j := 0; j < n; j++ {
		wg.Add(1)
		go RunWorker(ctx, taskChan, errCount, &wg)
	}

	if m > 0 {
		for {
			if int(*errCount) >= m {
				return ErrErrorsLimitExceeded
			}
			if len(taskChan) == 0 {
				break
			}
		}

	}
	wg.Wait()

	return nil
}

func RunWorker(ctx context.Context, tasks chan Task, errCounter *int32, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case task, ok := <-tasks:
			if !ok {
				return
			}

			err := task()
			if err != nil {
				atomic.AddInt32(errCounter, 1)
			}
		}
	}
}
