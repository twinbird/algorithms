package main

import (
	"testing"
)

var data []int
var expect []int
var maxValue int

func init() {
	data = []int{0, 1, 2, 5, 0, 0, 0, 5}
	expect = []int{0, 0, 0, 0, 1, 2, 5, 5}
	maxValue = 5
}

func TestCountingSort(t *testing.T) {
	countingSort(data, maxValue)

	for i, v := range expect {
		if data[i] != v {
			t.Fatalf("unexpected value index of %d\nexpect: %v\nactual: %v", i, expect, data)
		}
	}
}
