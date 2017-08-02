package main

import (
	"fmt"
)

func main() {
	data := []int{8, 4, 3, 7, 6, 5, 2, 1}

	fmt.Println(data)
	insertionSort(data)
	fmt.Println(data)
}

func insertionSort(data []int) {
	for i := 1; i < len(data); i++ {
		insert(data, i, data[i])
	}
}

func insert(data []int, pos int, val int) {
	var i int
	for i = pos - 1; 0 <= i && val < data[i]; i-- {
		data[i+1] = data[i]
	}
	data[i+1] = val
}
