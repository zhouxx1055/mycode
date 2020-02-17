package main

import (
	"log"
	"sync"
)

/**
2020/02/15 21:07:35 [0] is running
2020/02/15 21:07:35 [0] is done
2020/02/15 21:07:35 [1] is running
2020/02/15 21:07:35 [1] is done
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [semacquire]:
sync.runtime_Semacquire(0xc000062088)
	D:/go/src/runtime/sema.go:56 +0x40
sync.(*WaitGroup).Wait(0xc000062080)
	D:/go/src/sync/waitgroup.go:130 +0x6c
main.main()
	D:/ubcode/src/mycode/channel/case_deadlock1.go:21 +0x106

原因是 goroutine 得到的 "WaitGroup" 变量是 var wg WaitGroup 的一份拷贝值，即 doIt() 传参只传值。
所以哪怕在每个 goroutine 中都调用了 wg.Done()， 主程序中的 wg 变量并不会受到影响。

*/
// 等待所有 goroutine 执行完毕
// 进入死锁
func main() {
	var wg sync.WaitGroup
	done := make(chan struct{})

	workerCount := 2
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go doIt(i, done, wg)
	}

	close(done)
	wg.Wait()
	log.Println("all done!")
}

func doIt(workerID int, done <-chan struct{}, wg sync.WaitGroup) {
	log.Printf("[%v] is running\n", workerID)
	defer wg.Done()
	<-done
	log.Printf("[%v] is done\n", workerID)
}
