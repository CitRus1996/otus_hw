package hw05parallelexecution

import (
	"context"
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	wg := new(sync.WaitGroup)
	taskChan := make(chan Task, len(tasks))
	errChan := make(chan error, len(tasks))
	ctx := context.Background()
	for i := range tasks {
		taskChan <- tasks[i]
	}

	wg.Add(n)
	for j := 0; j < n; j++ {
		go RunWorker(ctx, taskChan, errChan, wg, m)
	}
	wg.Wait()

	if m > 0 {
		if len(errChan) >= m {
			return ErrErrorsLimitExceeded
		}
	}

	return nil
}

func RunWorker(ctx context.Context, tasks chan Task, errChan chan error, wg *sync.WaitGroup, m int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case task := <-tasks:
			err := task()
			if m > 0 {
				if err != nil {
					errChan <- err
					if len(errChan) >= m {
						ctx.Done()
						return
					}
				}
			}
		default:
			return
		}
	}
}
