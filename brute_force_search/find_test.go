package main

import (
	"testing"
)

func TestFindFound(t *testing.T) {
	idx := find(target, "commodo")
	if idx != 213 {
		t.Errorf("Expect:230, Actual:%d", idx)
	}
}

func TestFindNotFound(t *testing.T) {
	idx := find(target, "Hello")
	if idx != -1 {
		t.Errorf("Expect:-1, Actual:%d", idx)
	}
}
