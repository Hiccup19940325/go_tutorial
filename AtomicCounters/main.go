package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var ops atomic.Uint64
	// counter := 0

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				ops.Add(1)
				// counter++
			}

			wg.Done()
		}()
	}

	wg.Wait()
	// fmt.Println("ops:", counter)
	fmt.Println("ops:", ops.Load())
}
