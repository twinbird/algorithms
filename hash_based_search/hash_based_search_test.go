package main

import "testing"

func TestHashBasedSearch(t *testing.T) {
	data := []int{1, 4, 8, 9, 11, 15, 17}
	target := 11
	expect := true

	exist := hashBasedSearch(data, target)

	if exist != expect {
		t.Errorf("search result isn't expected.\nExpect:%t, Actual:%t\n",
			expect,
			exist)
	}
}

func TestHashBasedSearchNotFound(t *testing.T) {
	data := []int{1, 4, 8, 9, 11, 15, 17}
	target := 999
	expect := false

	exist := hashBasedSearch(data, target)

	if exist != expect {
		t.Errorf("search result isn't expected.\nExpect:%t, Actual:%t\n",
			expect,
			exist)
	}

}
