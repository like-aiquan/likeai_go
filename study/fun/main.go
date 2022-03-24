package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

/**
 * @author by chenaiquan<like.aiquan@qq.com> create on 2021/7/29 22.51
 */
func main() {
	//fmt.Println(print("hello", "world"))
	//fmt.Println(print("hello"))
	//fmt.Println(print())
	fmt.Println(count("how, do. you? 'dasda' ,... do? "))
	fmt.Println(add(1, 2))
	res, err := handlerError("aaa.text")
	fmt.Println(res)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hello("tom"))
	fmt.Println(get(6))
}

// catch
func get(index int) (ret int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Some error happened!", r)
			ret = -1
		}
	}()
	arr := []int{2, 3, 4}
	return arr[index]
}

func hello(name string) error {
	if len(name) == 0 {
		return errors.New("error: name is null")
	}
	fmt.Println("Hello,", name)
	return nil
}

func handlerError(res string) (string, error) {
	_, err := os.Open(res)
	return res, err
}

// 统计一个句子中每个单词出现的次数
func count(str string) (ret map[string]int) {
	str = strings.ReplaceAll(str, "\\pP", "")
	fmt.Println(str)
	keys := strings.Split(str, " ")
	ret = make(map[string]int, len(keys))
	for _, key := range keys {
		ret[key] += 1
	}
	return ret
}

func sum(x int, y int) (int, int) {
	return x + y, 1
}

func add(x int, y int) (z int) {
	z = x + y
	return
}

func print(str ...string) (string, string) {
	if len(str) > 1 {
		return str[0], str[1]
	} else if len(str) > 0 {
		return str[0], "nil"
	}
	return "nil", "nil"
}
