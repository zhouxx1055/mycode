package main

import (
	"fmt"
	"time"
)

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	// 斐波那契函数
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func main() {
	go spinner(100 * time.Millisecond) //开启一个goroutine
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFib(%d) = %d\n", n, fibN)
}
