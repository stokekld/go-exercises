package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mux sync.Mutex

	counter := 0
	gs := 100

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			mux.Lock()
			v := counter
			v++
			counter = v
			mux.Unlock()

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(counter)
}
