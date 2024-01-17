package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ExampleCycle() {
	// ExampleCyclePrimes()
	// ExampleCycleFiles()
	ExampleCycleFilesAnon()
}

// Дан список чисел, нужно вывести все простые числа из этого списка
func ExampleCyclePrimes() {
	nums := []int{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	response := []int{}

NUMS_LOOP:
	for _, num := range nums {
		if num < 2 {
			continue
		}
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				continue NUMS_LOOP
			}
		}
		response = append(response, num)
	}

	fmt.Println(response)
}

// Дан список файлов, для существующих файлов вывести название и заголовок
// Заголовок это часть файла до первой пустой строки
func ExampleCycleFiles() {
	fileNames := []string{"cycle.txt", "not_exists.go", "cycle.go"}
	for _, fName := range fileNames {
		fd, err := os.OpenFile(fName, os.O_RDONLY, 0666)
		if err != nil {
			continue
		}
		defer func(fName string, fd *os.File) {
			fd.Close()
			fmt.Println("Defer", fName)
		}(fName, fd)

		var sb strings.Builder
		reader := bufio.NewReader(fd)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println(fName, "ReadLine error", err)
				break
			}
			line = strings.TrimSpace(line)
			if line == "" {
				break
			}
			sb.WriteString(line)
			sb.WriteRune('\n')
		}
		fmt.Println(fName)
		fmt.Println("-----------------")
		fmt.Println(sb.String())
	}
}

func ExampleCycleFilesAnon() {
	fileNames := []string{"cycle.txt", "not_exists.go", "cycle.go"}
	for _, fName := range fileNames {
		func() { // Добавилась эта строка
			fd, err := os.OpenFile(fName, os.O_RDONLY, 0666)
			if err != nil {
				return // continue заменил на return
			}
			defer func(fName string, fd *os.File) {
				fd.Close()
				fmt.Println("Defer", fName)
			}(fName, fd)

			var sb strings.Builder
			reader := bufio.NewReader(fd)
			for {
				line, err := reader.ReadString('\n')
				if err != nil {
					fmt.Println(fName, "ReadLine error", err)
					return
				}
				line = strings.TrimSpace(line)
				if line == "" {
					break
				}
				sb.WriteString(line)
				sb.WriteRune('\n')
			}
			fmt.Println(fName)
			fmt.Println("-----------------")
			fmt.Println(sb.String())
		}() // Добавилась эта строка
	}
}
