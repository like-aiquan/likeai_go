// Package channel by chenaiquan<like.aiquan@qq.com> create on 2021/9/29 21.35
package channel

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

// 创建大小为 10 的缓冲信道
var ch1 = make(chan string, 10)

func download(url string) {
	fmt.Println("start to download", url)
	time.Sleep(time.Second)
	// 将 url 发送给信道
	ch1 <- url
}

func Test() {
	for i := 0; i < 3; i++ {
		go download("a.com/" + strconv.Itoa(i))
	}
	for i := 0; i < 3; i++ {
		msg := <-ch1 // 等待信道返回消息。
		fmt.Println("finish", msg)
	}
	fmt.Println("Done!")
}

var ch2 chan struct{} = make(chan struct{}, 2)

func foo() {
	ch2 <- struct{}{}
	log.Println("foo() 000")
	ch2 <- struct{}{}
	log.Println("foo() 111")
	time.Sleep(5 * time.Second)
	log.Println("foo() 222")
	close(ch2)
	log.Println("foo() 333")
}

func TestEmptyStruct() {
	var b struct{}
	log.Println("main() 111")
	go foo()
	log.Println("main() 222")
	a, ok := <-ch2
	log.Println("main() 333", a, ok)
	b, ok = <-ch2
	log.Println("main() 444", b, ok)
	// ch 此刻已经关闭
	// 从一个被close的channel中接收数据不会被阻塞，而是立即返回，接收完已发送的数据后会返回元素类型的零值(zero value)。
	// 从一个nil channel中接收数据会一直被block。（往nil channel中发送数据会一致被阻塞着。）
	c, ok := <-ch2
	// c : struct{}    ok = false
	log.Println("main() 555", c, ok)
}
