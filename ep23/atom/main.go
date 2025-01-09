package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// counter := 0
	var counter atomic.Uint32
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			for j := 0; j < 1000; j++ {
				// counter++
				counter.Add(1)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	// fmt.Println("Counter : ", counter)
	fmt.Println("Counter : ", counter.Load())
}
