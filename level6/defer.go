package main

import "fmt"

func main() {
	defer foo()
	fmt.Println("Hello world")
}

func foo() {
	fmt.Println("Hello foo")
}
