package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	fmt.Println("Hello world")

	wg.Add(2)

	go func() {
		fmt.Println("Hello from One")
		wg.Done()
	}()
	go func() {
		fmt.Println("Hello from Two")
		wg.Done()
	}()

	wg.Wait()
}
