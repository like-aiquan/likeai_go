// Package sort by chenaiquan<like.aiquan@qq.com> create on 2021/11/12 00.33
package sort

func QuickSort(arr []int, left int, right int) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}
	right = right - 1
	if left < right {
		p := partition(arr, left, right)
		QuickSort(arr, 0, p[0])
		QuickSort(arr, p[1], right)
	}
	return arr
}

func partition(arr []int, left int, right int) []int {
	less := left - 1
	more := right + 1
	flag := arr[right]
	for left < more {
		if arr[left] < flag {
			less = less + 1
			swap(arr, left, less)
			left = left + 1
		} else if arr[left] > flag {
			more = more - 1
			swap(arr, left, more)
		} else {
			left++
		}
	}
	return []int{less, more}
}
