package main

import (
	"fmt"
)

func main() {

	var x []int
	y := []int{123, 34, 8756, 23, 8, 4, 3}

	x = append(x, 8, 3, 1, 7)

	y = append(y, x...)

	fmt.Println(y)
	fmt.Println(y[2:5])

}
