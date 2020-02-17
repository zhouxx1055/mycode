package main

import (
	"fmt"
	"log"
	"time"
)

/**
有关 channel 的关闭，你需要注意以下事项:

关闭一个未初始化(nil) 的 channel 会产生 panic
重复关闭同一个 channel 会产生 panic
向一个已关闭的 channel 中发送消息会产生 panic
从已关闭的 channel 读取消息不会产生 panic，且能读出 channel 中还未被读取的消息，若消息均已读出，则会读到类型的零值。从一个已关闭的 channel 中读取消息永远不会阻塞，并且会返回一个为 false 的 ok-idiom，可以用它来判断 channel 是否关闭
关闭 channel 会产生一个广播机制，所有向 channel 读取消息的 goroutine 都会收到消息

作者：不智鱼
链接：https://www.jianshu.com/p/24ede9e90490
*/
func main() {
	//c := make(chan int)// 无缓冲   执行三次就结束
	c := make(chan int, 10)
	go send(c)
	go recv(c)
	time.Sleep(3 * time.Second)
	// 程序默认不等所有 goroutine 都执行完才退出，这点需要特别注意：
	// 常用解决办法：使用 "WaitGroup" 变量，它会让主程序等待所有 goroutine 执行完毕再退出。

	defer func() {
		close(c)
	}()

	//定义了一个缓冲大小为3的信道,可以放三个值
	var a chan int = make(chan int, 3)
	// 放三个值
	a <- 1
	a <- 2
	a <- 3
	// 也可以去三个字，有缓冲，不会造成阻塞态
	<-a
	<-a
	b := <-a
	fmt.Println("Hello go")
	fmt.Println(b)

}

//只能向chan里写数据
func send(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
		log.Println("send===>i:%d, len:%d", i, len(c))
	}
}

//只能取channel中的数据
func recv(c <-chan int) {
	for {
		if len(c) > 0 {
			for i := range c { //fatal error: all goroutines are asleep - deadlock!
				log.Println("recv===>i:%d, len:%d", i, len(c))
				time.Sleep(1 * time.Second)
			}
		} else {
			time.Sleep(1 * time.Second)
		}
	}
}
