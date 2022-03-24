// Package main by chenaiquan<like.aiquan@qq.com> create on 2021/10/21 13.29
package main

import (
	"fmt"
	"github.com/likeai/study/io/file"
)

func main() {
	//directory := "/Users/chenaiquan/go/src/github.com/likeai/io/console"
	//file.AddAnnotationOnProperty(directory)
	//fmt.Println()
	//excelFileName := "file name"
	//oraList := file.ReadXlsx(excelFileName)
	//file.WritingXlsx(oraList)
	fmt.Println(file.CaseToSnickedName("AbcdEfg"))

	type A interface{}
	var a A
	var b any
	fmt.Println(a == nil)
	fmt.Println(b == nil)
}
