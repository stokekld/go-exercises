package main

import (
	"fmt"
)

func main() {
	x := [5]int{123, 4, 23, 6, 87}

	for i, v := range x {
		fmt.Println(i, v)
	}
}
