package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	var wg sync.WaitGroup
	var errCount int32
	taskChan := make(chan Task, len(tasks))

	for i := range tasks {
		taskChan <- tasks[i]
	}
	close(taskChan)

	for j := 0; j < n; j++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for task := range taskChan {
				if !isLimitExceed(m, &errCount) {
					taskErr := task()
					if taskErr != nil {
						atomic.AddInt32(&errCount, 1)
					}
				}
			}
		}()
	}

	wg.Wait()
	if isLimitExceed(m, &errCount) {
		return ErrErrorsLimitExceeded
	}
	return nil
}

func isLimitExceed(maxErr int, counter *int32) bool {
	if maxErr > 0 {
		if int(atomic.LoadInt32(counter)) >= maxErr {
			return true
		}
	}
	return false
}
