// Package main by chenaiquan<like.aiquan@qq.com> create on 2021/11/11 23.50
package main

import (
	"fmt"
	"github.com/likeai/study/structure/sort"
	"reflect"
)

func main() {
	arr := []int{1, 2, 43, 5, 5, 6, 7}
	arr = sort.QuickSort(arr, 0, len(arr))
	var compare = func(o1 interface{}, o2 interface{}) bool {
		return o1.(int) > o2.(int)
	}
	fmt.Println(reflect.TypeOf(compare))
	array := make([]interface{}, len(arr))
	heap := sort.NewDefault(compare, array)
	fmt.Println(heap.Arr)
}
