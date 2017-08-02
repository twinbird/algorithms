package main

import "testing"

func TestSequentialSearch(t *testing.T) {
	data := []int{7, 8, 2, 388, 2, 3, 8, 9, 3, 13, 58, 7, 8, 3}
	target := 3
	expectPos := 5

	p := sequentialSearch(data, target)

	if p != expectPos {
		t.Errorf("search result isn't expected.\nExpect:%d, Actual:%d\n",
			expectPos,
			p)
	}
}

func TestSequentialSearchNotFound(t *testing.T) {
	data := []int{7, 8, 2, 388, 2, 3, 8, 9, 3, 13, 58, 7, 8, 3}
	target := 999
	expectPos := -1

	p := sequentialSearch(data, target)

	if p != expectPos {
		t.Errorf("search result isn't expected.\nExpect:%d, Actual:%d\n",
			expectPos,
			p)
	}

}
