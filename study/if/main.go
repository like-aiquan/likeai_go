package main

import (
	"fmt"
	"math"
)

/**
 * @author by chenaiquan<like.aiquan@qq.com> create on 2021/7/28 20.31
 */
func main() {
	v, ok := myFunc(-1.0)
	fmt.Println(v, ok)
}

func myFunc(f float64) (v float64, ok bool) {
	if f < 0 {
		return // default 0, false
	}
	return math.Sqrt(f), true
}
