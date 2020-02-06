package main

import "fmt"

/**
go 函数
一、语法
声明语法：func 函数名 (参数列表) [(返回值列表)] {}

二、特点
2.1不支持重载（可以编写多个同名函数，只要它们拥有不同的形参与/或者不同的返回值）
2.2函数可以赋值给变量
2.3多返回值
2.4匿名函数

三、传递方式：
3.1值传递
 　　基本的数据类型是值传递
3.2地址传递（引用传递）
　　map、slice、chan、指针、interface默认以引用的方式传递

四、入参
   ...代表函数可变

五、出参
   返回值命名
*/

func main() {
	ab, bc, cd := mingmin(3, 4)
	fmt.Println(ab, bc, cd)
	ab, bc, cd = ismingmin(3, 4)
	fmt.Println(ab, bc, cd)

	x := min(1, 2, 3, 0)
	fmt.Printf("The minimum is：%d\n", x)
	slice := []int{7, 9, 3, 5, 10}
	x = min(slice...)
	fmt.Printf("The minimum in the slice is: %d", x)
}

func mingmin(x1, x2 int) (x3, x4, x5 int) { // 返回值命名
	x3 = x1 + x2
	x4 = x1 * x2
	x5 = x1 - x2
	return
}

func ismingmin(x1, x2 int) (x3, x4, x5 int) {
	return x1 * x2, x1 + x2, x1 - x2
}

func min(s ...int) int { // ...代表函数可变
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}
