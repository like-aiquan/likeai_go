package main

import "fmt"

/**
 * @author by chenaiquan<like.aiquan@qq.com> create on 2021/7/28 11.21
 */
func main() {
	// 指向new的指针地址
	p := new(string)
	// 获取指针指向地址所指向的值  并修改
	*p = "hello pointer"
	// 打印指针指向地址指向的值
	fmt.Println(*p)

	str := "aaa"
	// 指向str的地址
	p = &str
	// 获取str指向的地址的值
	fmt.Println(*p)
	// 修改指向地址的值
	*p = "bbb"
	// 打印  str值发生变化  说明 *p 修改了str指向的值
	fmt.Println(str)

	p = nil
	//fmt.Println(p)
	// nil 指针再次反引用会报错
	//*p = "ccc"
	// 指向nil后拿*p的值也会抛错
	n, err := fmt.Printf(*p)
	if err != nil {
		fmt.Println(n, err)
	}

}
