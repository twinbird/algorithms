package main

import (
	"fmt"
	"os"
	"strconv"
)

func usage() {
	fmt.Println("Usage: fibonacci [number]")
	os.Exit(0)
}

func main() {
	if len(os.Args) != 2 {
		usage()
	}
	var num int
	var err error
	if num, err = strconv.Atoi(os.Args[1]); err != nil {
		usage()
	}
	fmt.Println(fibonacci(num))
}

func fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
