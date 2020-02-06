package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	//func1()
	//func2()
	//func3()
	func4()
}

func func1() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			fmt.Println(i)
			/**
			这种现象的原因在于闭包共享外部的变量i，
			注意到，每次调用go就会启动一个goroutine，这需要一定时间；
			但是，启动的goroutine与循环变量递增不是在同一个goroutine，可以把i认为处于主goroutine中。
			启动一个goroutine的速度远小于循环执行的速度，所以即使是第一个goroutine刚起启动时，外层的循环也执行到了最后一步了。
			由于所有的goroutine共享i，而且这个i会在最后一个使用它的goroutine结束后被销毁，所以最后的输出结果都是最后一步的i==5。
			*/
			wg.Done()
		}()
	}
	wg.Wait()
}

/*
输出结果：
5
5
5
5
5
*/

func func2() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
		time.Sleep(1 * time.Second) // 设置时间延时1秒
		// 每一步循环至少间隔一秒，而这一秒的时间足够启动一个goroutine了，因此这样可以输出正确的结果。
	}
	wg.Wait()
}

/*
输出结果：
0
1
2
3
4
*/

func func3() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i) // 共享的环境变量作为函数参数传递
	}
	wg.Wait()
}

/*
输出:
4
0
3
1
2
*/

func func4() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		i := i // 注意这里的同名变量覆盖  使用同名的变量保留当前的状态
		// 同名的变量i作为内部的局部变量，覆盖了原来循环中的i，
		// 此时闭包中的变量不在是共享外循环的i，而是都有各自的内部同名变量i，
		// 赋值过程发生于循环goroutine，因此保证了独立。
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}
	wg.Wait()
}

/*
输出结果：
4
2
0
3
1
结果顺序原因同1
*/
