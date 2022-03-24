// Package memory by chenaiquan<like.aiquan@qq.com> create on 2021/10/22 15.40
package memory

import (
	"fmt"
	"unsafe"
)

type Demo struct {
	// empty struct
	m struct{}
	// int 8  only one byte
	x int8
}

type Demo2 struct {
	x int8
	m struct{}
}

var d1 Demo
var d2 Demo2

func Test() {
	fmt.Println(unsafe.Sizeof(d1))
	// int 8 类型只需要对其一个字节
	fmt.Println(unsafe.Alignof(d1.x))
	// 空结构体如果不在结构体的最后不需要对齐 但此处显示他需要对齐一个 bit
	fmt.Println(unsafe.Alignof(d1.m))

	fmt.Println(unsafe.Sizeof(d2))
	// int 8 类型只需要对其一个字节
	fmt.Println(unsafe.Alignof(d2.x))
	// 空结构体在结构体定义的最后  需要对齐 避免内存泄漏
	fmt.Println(unsafe.Alignof(d2.m))

	fmt.Println(unsafe.Sizeof(d2.m))

}
