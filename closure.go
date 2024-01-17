package main

import (
	"fmt"
	"html"
	"net/http"
	"slices"
	"sort"
	"strings"
)

func ExampleClosure() {
	ExampleClosureSimple()
	// ExampleClosureLikeObject()
	// ExampleClosureHTTPServer()
	// ExampleClosureSort()
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

	fmt.Println("Two")
	inc2, dec2, get2 := CreateClosureLikeObject()
	fmt.Println(get2())
	inc2(30)
	fmt.Println(get2())
	dec2(20)
	fmt.Println(get2())

	fmt.Println("1 not changed", get1())
}

func ExampleClosureHTTPServer() {
	// https://pkg.go.dev/net/http@go1.21.6#hdr-Servers
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
}

func ExampleClosureSort() {
	SortCaseIgnore()
}

func SortCaseIgnore() {
	// Дан слайс строк, отсортировать его без учёта регистра
	slice := []string{
		"АНТОШКА", "АНтошка", "пойдём", "Копать", "картошку",
		"антошкА", "антошКА", "пойдём", "копать", "карто-ошку",
		"тили", "тили", "трали", "вали",
	}

	// Old way
	unstable := make([]string, len(slice))
	copy(unstable, slice)
	// Unstable sort with Less
	sort.Slice(unstable, func(i, j int) bool {
		return strings.ToLower(unstable[i]) < strings.ToLower(unstable[j])
	})
	fmt.Println("Unstable sort", unstable)

	// Go 1.21
	stable := slices.Clone(slice)
	// Stable sort with Compare
	slices.SortStableFunc(stable, func(a, b string) int {
		aLower, bLower := strings.ToLower(a), strings.ToLower(b)
		if aLower == bLower {
			return 0
		}
		if aLower < bLower {
			return -1
		}
		return 1
		//return cmp.Compare(strings.ToLower(a), strings.ToLower(b))
	})
	fmt.Println("Stable sort", stable)
}
