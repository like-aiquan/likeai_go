// Package sync by chenaiquan<like.aiquan@qq.com> create on 2021/9/29 17.11
package sync

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

// Download 协程的一些东西 wg 协程组
func Download(url string, wg *sync.WaitGroup) {
	fmt.Println("start download", url)
	// 休眠一秒
	time.Sleep(time.Second)
	// 协程数 -1  当协程都执行完了 waitGroup 数量应该是0
	wg.Done()
}

func Test() {
	fmt.Println(time.Now().Format("15:04:05"))
	for i := 0; i < 3; i++ {
		// 新增协程数协程
		wg.Add(1)

		// 开启协程
		// 跨包方法调用  wg 不是引用类型 指针传递 否则传递的是引用会产生死锁
		go Download("a.com/"+strconv.Itoa(i), &wg)
	}
	// 等待所有协程结束任务
	wg.Wait()
	fmt.Println("success!")
}
