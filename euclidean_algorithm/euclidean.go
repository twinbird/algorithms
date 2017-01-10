package main

import (
	"fmt"
	"os"
	"strconv"
)

func usage() {
	fmt.Println("Usage: euclidean [number] [number]")
	os.Exit(0)
}

func main() {
	if len(os.Args) != 3 {
		usage()
	}
	var a, b int
	var err error
	if a, err = strconv.Atoi(os.Args[1]); err != nil {
		usage()
	}
	if b, err = strconv.Atoi(os.Args[2]); err != nil {
		usage()
	}

	if a < b {
		tmp := a
		a = b
		b = tmp
	}

	euclidean(a, b)
}

func euclidean(a, b int) {
	m := a % b
	if m == 0 {
		fmt.Println(b)
		return
	}
	euclidean(b, m)
}
