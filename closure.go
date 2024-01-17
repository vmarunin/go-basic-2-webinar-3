package main

import "fmt"

func ExampleClosure() {
	// ExampleClosureSimple()
	ExampleClosureLikeObject()
}

func ExampleClosureSimple() {
	counter := 0
	increment := func(inc int) {
		counter += inc
	}
	fmt.Println("Before", counter)
	increment(3)
	fmt.Println("After", counter)
}

func CreateClosureLikeObject() (inc, dec func(v int), get func() int) {
	var counter int = 0
	incFunc := func(v int) {
		counter += v
	}
	decFunc := func(v int) {
		counter -= v
	}
	getFunc := func() int {
		return counter
	}
	return incFunc, decFunc, getFunc
}

func ExampleClosureLikeObject() {
	fmt.Println("One")
	inc1, dec1, get1 := CreateClosureLikeObject()
	fmt.Println(get1())
	inc1(3)
	fmt.Println(get1())
	dec1(2)
	fmt.Println(get1())

	fmt.Println("One")
	inc2, dec2, get2 := CreateClosureLikeObject()
	fmt.Println(get2())
	inc2(30)
	fmt.Println(get2())
	dec2(20)
	fmt.Println(get2())

	fmt.Println("1 not changed", get1())
}
