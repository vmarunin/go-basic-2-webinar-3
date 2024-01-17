package main

import "fmt"

func RegularFunction(in string) string {
	return "I am regular " + in
}

var globalAnonymousFunction = func(in string) string {
	return "I am global anonymous " + in
}

func ExampleSignature() {
	anonFunction := func(in string) string {
		return "I am anonymous " + in
	}
	myFynction := anonFunction
	fmt.Println(myFynction("test1"))

	myFynction = RegularFunction
	fmt.Println(myFynction("test2"))

	myFynction = func(in string) string {
		return "I am another anonymous " + in
	}
	fmt.Println(myFynction("test3"))

	myFynction = globalAnonymousFunction
	fmt.Println(myFynction("test4"))

	var recursiveAnonymousFunction func(counter int)
	recursiveAnonymousFunction = func(counter int) {
		fmt.Println("I am recursive anonymous function", counter)
		if counter == 0 {
			return
		}
		recursiveAnonymousFunction(counter - 1)
	}
	recursiveAnonymousFunction(5)

	// CallRightNow
	func(in string) {
		fmt.Println("I am right now", in)
	}("test5")
}
