package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * @author by chenaiquan<like.aiquan@qq.com> create on 2021/7/25 14.28
 */
func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d, ", arr[i])
	}
	fmt.Println()
	// 切片
	section := arr[3:6]
	// 切片是引用类型
	fmt.Println(section)
	arr[5] = 10
	fmt.Println(section)
	// 切片对象  类似于 [切点index, 切片长度, 切片容量]   len 标识切片长度   cap标识容量
	// 切片的容量是从底层数组开始算的, 例如下 len 3, cap 6   cap表示的是从4开始到数组元素末尾的容量 = 6
	fmt.Printf("len %d, cap %d\n", len(section), cap(section))
	// 如下切片
	section3 := arr[:5]
	// 切片序列
	fmt.Println(section3)
	// len 5, cap 9    5 表示切片的长度,  9代表底层数组的容量   9 = 切点index --> 最后数组一个位置的长度
	fmt.Printf("len %d, cap %d\n", len(section3), cap(section3))
	// 切片再切片
	section1 := section[2:]
	fmt.Println(section1)
	// 直接定义切片 [...]不定长数组   []切片
	section2 := []string{"中国", "北京", "天安门"}
	fmt.Println(section2)

	// make 函数创建切片
	slice := make([]int, 5, 10)
	// slice [0 0 0 0 0], len 5, cap 10
	fmt.Printf("slice %v, len %d, cap %d\n", slice, len(slice), cap(slice))

	// nil
	var slice1 []int
	// nil 类似于null  表示引用为null  nil的切片长度一定是0
	fmt.Printf("%v, len %d, cap %d\n", slice1 == nil, len(slice1), cap(slice1))
	// 长度为0 的切片不一定是nil
	slice2 := []int{}
	fmt.Printf("%v, len %d, cap %d\n", slice2 == nil, len(slice2), cap(slice2))
	slice3 := make([]int, 0)
	fmt.Printf("%v, len %d, cap %d\n", slice3 == nil, len(slice3), cap(slice3))

	slice4 := arr[3:6]
	slice5 := slice4
	// 指向同一个底层数组
	fmt.Println(slice5, slice4)
	// 赋值
	slice4[0] = 1000
	fmt.Println(slice5, slice4, arr)

	for i, v := range slice4 {
		fmt.Printf("%d %v\n", i, v)
	}

	// 切片追加
	slice6 := []string{"北", "上", "广"}
	fmt.Printf("%v, len %d, cap %d\n", slice6, len(slice6), cap(slice6))
	slice6 = append(slice6, "深")
	// 存不下就扩容
	fmt.Printf("%v, len %d, cap %d\n", slice6, len(slice6), cap(slice6))
	// 切片追加切片
	slice7 := []string{"西安", "杭州", "成都"}
	// ... 标记拆分切片内元素
	slice6 = append(slice6, slice7...)
	fmt.Printf("%v, len %d, cap %d\n", slice6, len(slice6), cap(slice6))

	str := "i am wa wa wa"
	fields := strings.Fields(str)
	for i, v := range fields {
		fmt.Printf("%d, %v\t", i, v)
	}
	fmt.Println()
	fmt.Printf("%d \n", strings.Count(str, "wa"))

	// 获取运行环境的int字节   println是异步的
	println(strconv.IntSize)
	f := 1.232
	// 浮点型转换成string类型
	fmt.Printf("%s , %T \n", strconv.FormatFloat(f, 'f', 4, 64), strconv.FormatFloat(f, 'f', 4, 64))
	formatFloat := strconv.FormatFloat(f, 'f', 4, 64)
	formatFloat = "2123121a"
	f, err := strconv.ParseFloat(formatFloat, 64)
	if err != nil {
		fmt.Printf("parse error!")
	}

	fmt.Println(f)
	at, err := strconv.Atoi("221")
	fmt.Println(at)
	if err != nil {
		fmt.Printf("parse error!")
	}
	itoa := strconv.Itoa(12312)
	fmt.Println(itoa)

}
