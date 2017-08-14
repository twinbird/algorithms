package main

import (
	"fmt"
)

const target = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."

func find(data string, keyword string) int {
	keylen := len(keyword)
	datalen := len(data)

	for i := range data {
		if (i + keylen) > datalen {
			return -1
		}
		if match(data[i:i+keylen], keyword) == true {
			return i
		}
	}
	return -1
}

func match(data string, keyword string) bool {
	for i := range keyword {
		if data[i] != keyword[i] {
			return false
		}
	}
	return true
}

func main() {
	idx := find(target, "commodo")
	if idx < 0 {
		fmt.Println("Not found")
	} else {
		fmt.Printf("Found at %d\n", idx)
	}
}
