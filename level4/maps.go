package main

import (
	"fmt"
)

func main() {

	x := map[string]int{
		"Jesus": 32,
	}

	x["James"] = 33
	x["Penny"] = 34

	for k, v := range x {
		fmt.Println(k, v)
	}

	delete(x, "Jesus")

	fmt.Println(x)
}
