// Package _struct by chenaiquan<like.aiquan@qq.com> create on 2021/9/29 16.41
package _struct

type Student struct {
	name string
	age  int
}

func (stu *Student) GetAge() int {
	return stu.age
}
