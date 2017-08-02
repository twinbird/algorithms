package main

import (
	"fmt"
)

func main() {
	data := []int{1, 4, 8, 9, 11, 15, 17}
	target := 11

	ok := hashBasedSearch(data, target)
	if ok == false {
		fmt.Println("%d is not found", target)
	} else {
		fmt.Printf("%d is found.", target)
	}
}

func hashBasedSearch(data []int, target int) bool {
	t := newHashTable(data)

	h := hash(target)
	list := t[h]

	for _, v := range list {
		if v == target {
			return true
		}
	}
	return false
}

func newHashTable(data []int) [][]int {
	t := make([][]int, 3)

	for i := 0; i < len(data); i++ {
		h := hash(data[i])
		if t[h] == nil {
			t[h] = make([]int, 0)
		}
		t[h] = append(t[h], data[i])
	}
	return t
}

func hash(x int) int {
	return x % 3
}
