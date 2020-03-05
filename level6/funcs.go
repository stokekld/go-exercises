package main

import "fmt"

func main() {

	x := foo()
	n, s := bar()

	fmt.Println(x)

	fmt.Println(n, s)
}

func foo() int {
	return 42
}

func bar() (int, string) {
	return 1984, "Big brother"
}
