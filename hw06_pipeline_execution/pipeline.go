package hw06pipelineexecution

import "sync"

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	var wg sync.WaitGroup
	pool := make([]Bi, len(stages)+1)
	for i := range pool {
		pool[i] = make(Bi, 10)
	}
	wg.Add(1)
	go func() {
		defer wg.Done()

		for data := range in {
			pool[0] <- data
		}

		close(pool[0])
	}()

	for i := 0; i < len(stages); i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()

			for data := range stages[j](pool[j]) {
				pool[j+1] <- data
			}

			close(pool[j+1])
		}(i)
	}

	wg.Wait()
	return pool[len(stages)]
}
