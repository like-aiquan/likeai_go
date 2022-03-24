// Package _interface @author by chenaiquan<like.aiquan@qq.com> create on 2021/9/29 15.50
package main

import (
	"fmt"
	"strconv"
)

type Person interface {
	// getName 小写代表 private
	getName() string
	// toString 自定义
	toString() string
	// GetAge 大写代表 public
	GetAge() int
}

type Man interface {
	// GetAge interface 里面不能有私有方法 不然其他的包不能实现这个方法
	GetAge() int
}

// Woman 空接口代表可以被全部类型实现, 即它是全部类型的父类
type Woman interface{}

type Student struct {
	name string
	age  int
}

func (stu *Student) getName() string {
	return stu.name
}

// GetAge 同时实现了两个 interface, golang 支持多实现
func (stu *Student) GetAge() int {
	return stu.age
}

func (stu *Student) toString() string {
	return stu.name + " " + strconv.Itoa(stu.age) + "years old"
}

func main() {
	var p Person = &Student{
		name: "Tom",
		age:  18,
	}
	typeMap := map[string]Woman{}
	typeMap["int"] = 120
	typeMap["string"] = "string"
	typeMap["Student"] = p.toString()
	fmt.Println(typeMap)
	fmt.Println(p.getName())
}
