package main

import (
	"bufio"
	"fmt"
	"os"
)

func ExampleDefer() {
	ExampleDeferStack()
	// ExampleDeferRegular()
	// ExampleDeferParams()
	// ExampleDeferReturn()
	// ExampleDeferPanic()
	// ExampleDeferFile()
}

func ExampleDeferStack() {
	defer fmt.Println("defer 1")

	fmt.Println("function 1")
	fmt.Println("function 2")

	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")

	fmt.Println("function 3")
}

func deferHelpeer(s string) {
	fmt.Println("Defer helper", s)
}
func ExampleDeferRegular() {
	defer deferHelpeer("1")
	fmt.Println("Regular 1")
}

func ExampleDeferParams() {
	counter := 0
	deferHelper := func(v int) {
		fmt.Println("deferHelper", "counter", counter, "v", v)
	}
	defer deferHelper(1)
	counter = 1
	defer deferHelper(2)
	fmt.Println("ExampleDeferParams")
}

func returnHelpeer(s string) (ret string) {
	defer func() {
		if s == "defer" {
			ret = "This is SPARTA!!!"
		}
	}()
	return s
}
func ExampleDeferReturn() {
	fmt.Println("regular", returnHelpeer("regular"))
	fmt.Println("defer", returnHelpeer("defer"))
}

func helperDeferPanic() {
	// panic("I am panic")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered", err)
		}
	}()
	panic("I am panic")

	fmt.Println("After panic")
}
func ExampleDeferPanic() {
	fmt.Println("Before panic helper")
	helperDeferPanic()
	fmt.Println("After panic helper")
}

func ExampleDeferFile() {
	fd, err := os.OpenFile("./defer.go", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("OpenFile error", err)
		return
	}
	defer fd.Close()

	reader := bufio.NewReader(fd)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("ReadLine error", err)
		return
	}
	fmt.Println("ReadLine", line)
}
