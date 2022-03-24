package main

import "fmt"

/**
 * @author by chenaiquan<like.aiquan@qq.com> create on 2021/8/29 17.24
 */
func main() {
	s1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(s1)
	s2 := s1
	fmt.Println(s2)
	// 初始化内存 len 必须指定为3
	// 不然切片内没有任何内容, 虽然数组已经拷贝了
	s3 := make([]int, 5, 5)
	copy(s3, s2)
	s2[1] = 100
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}
