pakcage main

import (
	"fmt"
	"strconv"
	"os"
)

func usage() {
	fmt.Println("Usage: hanoi [disk number]")
	os.Exit(0)
}

func main() {
	if len(os.Args) != 2 {
		usage()
	}
	var disks int
	var err error
	if disks, err = strconv.Atoi(os.Args[1]) {
		usage()
	}
}

type tower struct {
}
