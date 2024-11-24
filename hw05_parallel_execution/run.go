package hw05parallelexecution

import (
	"errors"
	"fmt"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.

func Run(tasks []Task, n, m int) error {
	if n <= 0 {
		fmt.Println("Capacity can't be negative value")
		return nil
	}

	// Helper struct to arrange errors control.
	var errCtrl struct {
		mu sync.Mutex
		cnt int
	}

	var ch = make(chan Task, n)

	var wg sync.WaitGroup
	wg.Add(n)

	// Create n workers, which will read channel with tasks queue
	for i := 0; i < n; i++ {
		go func() {
			for {
				// Verify if errors counter exceeds the limit already.
				if errCtrl.cnt >= m && m > 0 {
					// Drain data in channel, otherwise we'll be unable to close it later.
					for {
						_, ok := <-ch
						if !ok {
							defer wg.Done()
							break
						}
					}
					return
				}

				t, ok := <-ch

				if !ok {
					defer wg.Done()
					return
				}

				res := t()
				// Increment total errors value, use mutex to access errors counter,
				// to avoid decremented counter when one thread overwrites another's value.
				// For sure could use atomic.AddIntNN as well.
				if res != nil {
					errCtrl.mu.Lock()
					errCtrl.cnt++
					errCtrl.mu.Unlock()
				}
			}
		}()
	}

	// put tasks into the queue via channel
	for _, t := range tasks {
		ch <- t
	}
	close(ch)
	wg.Wait()
	if errCtrl.cnt >= m && m > 0 {
		return ErrErrorsLimitExceeded
	}
	return nil
}
