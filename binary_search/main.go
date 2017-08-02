package main

import (
	"fmt"
)

func main() {
	data := []int{1, 4, 8, 9, 11, 15, 17}
	target := 11

	p := binarySearch(data, target)
	if p < 0 {
		fmt.Println("%d is not found", target)
	} else {
		fmt.Printf("%d is found at %d.\ndata:%v", target, p, data)
	}
}

func binarySearch(data []int, target int) int {
	low := 0
	high := len(data) - 1

	for low <= high {
		mid := (low + high) / 2
		if target == data[mid] {
			return mid
		} else if target < data[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}
