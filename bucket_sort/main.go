package main

import (
	"fmt"
)

func main() {
	data := []int{0, 1, 2, 5, 0, 0, 0, 5}

	fmt.Println(data)
	bucketSort(data)
	fmt.Println(data)
}

func bucketSort(data []int) {
	// 2次元のバケットを作成
	bucket := make([][]int, len(data))
	for i := 0; i < len(bucket); i++ {
		bucket[i] = make([]int, 0)
	}

	// ハッシュ値に合わせてバケットへ入れる
	for i := 0; i < len(data); i++ {
		k := hash(data[i])
		bucket[k] = append(bucket[k], data[i])
	}

	var idx int
	for i := 0; i < len(bucket); i++ {
		// バケット内でまず整列
		insertionSort(bucket[i])
		for k := 0; k < len(bucket[i]); k++ {
			// 整列済みバケットから順に取り出して入れる
			data[idx] = bucket[i][k]
			idx++
		}
	}
}

func hash(x int) int {
	return x / 3
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
