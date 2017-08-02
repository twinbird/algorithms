package main

import (
	"fmt"
)

func main() {
	data := []int{8, 4, 3, 7, 6, 5, 2, 1}

	fmt.Println(data)
	heapSort(data)
	fmt.Println(data)
}

func heapSort(data []int) {
	// ヒープを構築
	buildHeap(data)
	for i := len(data) - 1; i >= 1; i-- {
		// ヒープの根は最大値なので入れ替え
		data[0], data[i] = data[i], data[0]
		// ヒープ再構築
		heapify(data, 0, i)
	}
}

// 2分ヒープ木を構築
func buildHeap(data []int) {
	// 木の高さは2分木なのでデータ数/2になるはず
	// len(data)/2回再構築処理を行う
	for i := len(data)/2 - 1; i >= 0; i-- {
		heapify(data, i, len(data))
	}
}

// idxは木の階層, maxは配列の長さ.
// ヒープは上位が下位よりも必ず大きい事が保証された木.
//     idx
//    /   \
//  left  right
// と考えて部分木を構築しなおしていく.
func heapify(data []int, idx, max int) {
	leftIdx := 2*idx + 1
	rightIdx := 2*idx + 2

	var largestIdx int
	if leftIdx < max && data[leftIdx] > data[idx] {
		largestIdx = leftIdx
	} else {
		largestIdx = idx
	}
	if rightIdx < max && data[rightIdx] > data[largestIdx] {
		largestIdx = rightIdx
	}
	// 最大値が変わったという事は値交換と下位の部分木の再構築が必要
	if largestIdx != idx {
		data[idx], data[largestIdx] = data[largestIdx], data[idx]
		heapify(data, largestIdx, max)
	}
}
