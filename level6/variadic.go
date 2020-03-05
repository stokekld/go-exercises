package main

import "fmt"

func main() {

	ii := []int{1, 2, 3, 4, 5, 6, 7}
	x := foo(ii...)

	fmt.Println(x)
}

func foo(xi ...int) int {
	total := 0

	for _, v := range xi {
		total += v
	}

	return total
}
