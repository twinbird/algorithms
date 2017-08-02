package main

import "testing"

func TestBinarySearch(t *testing.T) {
	data := []int{1, 4, 8, 9, 11, 15, 17}
	target := 11
	expectPos := 4

	p := binarySearch(data, target)

	if p != expectPos {
		t.Errorf("search result isn't expected.\nExpect:%d, Actual:%d\n",
			expectPos,
			p)
	}
}

func TestBinarySearchNotFound(t *testing.T) {
	data := []int{1, 4, 8, 9, 11, 15, 17}
	target := 999
	expectPos := -1

	p := binarySearch(data, target)

	if p != expectPos {
		t.Errorf("search result isn't expected.\nExpect:%d, Actual:%d\n",
			expectPos,
			p)
	}

}
