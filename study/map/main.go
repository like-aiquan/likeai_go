package main

import (
	"fmt"
	"sort"
)

/**
 * @author by chenaiquan<like.aiquan@qq.com> create on 2021/8/22 15.22
 */
func main() {
	// map 必须初始化
	m1 := make(map[string]int, 1)
	m1["理想"] = 9000
	fmt.Println(m1)
	m1["李白"] = 200
	fmt.Println(m1)
	// 查询不到这个 value 默认 int = 0
	fmt.Println(m1["杜甫"])
	i, ok := m1["杜甫"]
	if !ok {
		fmt.Println("查无此key")
	} else {
		fmt.Println(i)
	}

	m2 := make(map[string]string, 1)
	// 没有这个key  一个空字符串
	fmt.Println(m2["杜甫"])

	for k, v := range m1 {
		fmt.Println(k, v)
	}

	for k := range m1 {
		fmt.Println(k)
	}

	for _, v := range m1 {
		fmt.Println(v)
	}

	delete(m1, "李白")
	fmt.Println(m1)
	// 不存在的key 不做任何操作
	delete(m1, "杜甫")
	fmt.Println(m1)

	m1["李白"] = 100
	m1["白居易"] = 200

	// 排序取map
	keys := make([]string, 0, 200)
	for k := range m1 {
		keys = append(keys, k)
	}
	// sort key
	sort.Strings(keys)
	fmt.Println(keys)
}
