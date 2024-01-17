package main

import "fmt"

func ExampleDefer() {
	ExampleDeferStack()
}

func ExampleDeferStack() {
	defer fmt.Println("defer 1")

	fmt.Println("function 1")
	fmt.Println("function 2")

	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")

	fmt.Println("function 3")
}
