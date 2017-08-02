package main

import (
	"testing"
)

var data []int
var expect []int

func init() {
	data = []int{8, 4, 3, 7, 6, 5, 2, 1}
	expect = []int{1, 2, 3, 4, 5, 6, 7, 8}
}

func TestQuickSort(t *testing.T) {
	quickSort(data)

	for i, v := range expect {
		if data[i] != v {
			t.Fatalf("unexpected value index of %d\nexpect: %v\nactual: %v", i, expect, data)
		}
	}
}
