package main

import (
	"fmt"
	"runtime"
)

var (
	name string
	age  int
	isOk bool
)

const sb = "const test"

const (
	n1 = 100
	n2
	n3
)
const (
	i1 = iota // this 0
	i2        // auto increment  this iota = 1
	i3        // ditto
)

func main() {
	name = "like.ai"
	age = 23
	isOk = true
	fmt.Println("hello world, ", runtime.Version())
	fmt.Printf("hello %s!", name)
	fmt.Println()
	fmt.Printf("%s is a %d years old boy!\n", name, age)
	fmt.Println(isOk)
	str := "qqai"
	fmt.Println(str)
_:
	s := "com"
	fmt.Println(s)
	fmt.Println(sb)
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(i1)
	fmt.Println(i2)
	fmt.Println(i3)

	s = "like ai"
	for i, v := range s {
		fmt.Printf("%d, %c, \n", i, v)
	}

	b := 0b10000111
	fmt.Printf("%d, %T", b, b)
}
