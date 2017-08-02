package main

import (
	"fmt"
)

func main() {
	data := []int{8, 4, 3, 7, 6, 5, 2, 1}

	fmt.Println(data)
	selectionSort(data)
	fmt.Println(data)
}

func selectionSort(data []int) {
	var minIdx int
	for i := 0; i < len(data); i++ {
		minIdx = i
		for k := i; k < len(data); k++ {
			if data[k] < data[minIdx] {
				minIdx = k
			}
		}
		data[i], data[minIdx] = data[minIdx], data[i]
	}
}
