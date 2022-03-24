package main

import "fmt"

/**
 * @author by chenaiquan<like.aiquan@qq.com> create on 2021/7/25 01.05
 */
func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", i, j, i*j)
		}
		fmt.Println()
	}

}
