package main

import (
	"fmt"
)

func main() {
	data := []int{7, 8, 2, 388, 2, 3, 8, 9, 3, 13, 58, 7, 8, 3}
	target := 3

	p := sequentialSearch(data, target)
	if p < 0 {
		fmt.Println("%d is not found", target)
	} else {
		fmt.Printf("%d is found at %d.\ndata:%v", target, p, data)
	}
}

func sequentialSearch(data []int, target int) int {
	for i := 0; i < len(data); i++ {
		if data[i] == target {
			return i
		}
	}
	return -1
}
