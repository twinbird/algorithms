package main

import (
	"fmt"
)

func main() {
	data := []int{8, 4, 3, 7, 6, 5, 2, 1}

	fmt.Println(data)
	quickSort(data)
	fmt.Println(data)
}

func quickSort(data []int) {
	quickSortSub(data, 0, len(data)-1)
}

func quickSortSub(data []int, leftIdx, rightIdx int) {
	if leftIdx >= rightIdx {
		return
	}
	p := partition(data, leftIdx, rightIdx)
	quickSortSub(data, leftIdx, p-1)
	quickSortSub(data, p+1, rightIdx)
}

// 指定区間で半分に分けて値を入れ替える
// 分けた地点のインデックスを返す
func partition(data []int, leftIdx, rightIdx int) int {
	// この区間の中間地点
	midIdx := (leftIdx + rightIdx - 1) / 2
	// ピボットに使う値のインデックス
	pivotIdx := med3(data, leftIdx, midIdx, rightIdx)
	// ピボットと右端を交換(退避)
	data[pivotIdx], data[rightIdx] = data[rightIdx], data[pivotIdx]
	// 値交換に使うインデックス
	// 半分に分けた左側からピボット値(右端に入れたもの)を基準に入れ替えていく
	storeIdx := leftIdx
	for i := leftIdx; i < rightIdx; i++ {
		if data[i] <= data[rightIdx] {
			// 値交換
			data[i], data[storeIdx] = data[storeIdx], data[i]
			storeIdx++
		}
	}
	// 退避したものを戻す
	data[storeIdx], data[rightIdx] = data[rightIdx], data[storeIdx]
	// 最後に値交換した場所が分割した場所
	return storeIdx
}

// 中央値のインデックスを返す
func med3(data []int, x, y, z int) int {
	if data[x] < data[y] {
		if data[y] < data[z] {
			return y
		} else if data[z] < data[x] {
			return x
		} else {
			return z
		}
	} else {
		if data[z] < data[y] {
			return y
		} else if data[x] < data[z] {
			return x
		} else {
			return z
		}
	}
}
