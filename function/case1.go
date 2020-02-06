package main

import "fmt"

// 高阶函数   ：函数作为参数或者返回值

func main() {
	f1()

	// 函数可以作为参数
	ff1 := f1
	f2(ff1)
}

func f1() {
	fmt.Println("function f1")
}

func f2(x func()) {
	x()
}
