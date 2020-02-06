package main

import "fmt"

// 高阶函数   ：函数作为参数或者返回值

func iMap(num []int, f func(a int) int) []int {
	var r []int

	for _, n := range num {
		r = append(r, f(n))
	}

	return r
}

func main() {
	num := []int{2, 4, 6, 8}

	fmt.Println("before: ", num)

	result := iMap(num, func(a int) int {
		return a * 2
	})

	fmt.Println("after: ", result)
}
