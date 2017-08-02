package main

import (
	"fmt"
)

func main() {
	data := []int{0, 1, 2, 5, 0, 0, 0, 5}
	maxValue := 5

	fmt.Println(data)
	countingSort(data, maxValue)
	fmt.Println(data)
}

func countingSort(data []int, maxVal int) {
	bucket := make([]int, maxVal+1)

	for i := 0; i < len(data); i++ {
		bucket[data[i]]++
	}

	var idx int
	for i := 0; i < maxVal+1; i++ {
		for bucket[i] > 0 {
			data[idx] = i
			idx++
			bucket[i]--
		}
	}
}
