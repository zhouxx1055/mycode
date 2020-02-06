package main

import "fmt"

// 闭包

func f11(f func()) {
	fmt.Println("this is f11")
	f()
}
func f22(x int, y int) {
	fmt.Println("this is f22")
	fmt.Println(x + y)
}

// 定义一个高阶函数 对f11进行包装
func f14(f func(int, int), x, y int) func() {
	tmp := func() {
		f(x, y)
	}
	return tmp
}

func f13() func() {
	tmp := func() {
		fmt.Println("this is f13")
	}
	return tmp
}

func main() {
	f := f13()
	f()

	f0 := f14(f22, 3, 4)
	f0()
}
