// Package sort by chenaiquan<like.aiquan@qq.com> create on 2021/11/11 23.48
package sort

func HeapSort(arr []int) []int {
	heapSize := len(arr)
	if arr == nil || heapSize < 2 {
		return arr
	}
	for i := 0; i < heapSize; i++ {
		heapInsert(arr, i)
	}
	heapSize = heapSize - 1
	swap(arr, 0, heapSize)
	for heapSize > 0 {
		heapify(arr, 0, heapSize)
		heapSize = heapSize - 1
		swap(arr, 0, heapSize)
	}
	return arr
}

func heapify(arr []int, index int, heapSize int) {
	left := index*2 + 1
	for left < heapSize {
		largestIndex := left
		if left+1 < heapSize && arr[left+1] > arr[left] {
			largestIndex = left + 1
		}
		if arr[largestIndex] > arr[index] {
			swap(arr, largestIndex, index)
		}
		index = largestIndex
		left = index*2 + 1
	}
}

func heapInsert(arr []int, i int) {
	for arr[i] > arr[(i-1)/2] {
		swap(arr, i, (i-1)/2)
		i = i / 2
	}
}

func swap(arr []int, i int, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}
