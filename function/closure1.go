package main

import "fmt"

/**
函数在Golang中是“一等公民”，因此关于函数的特性必须要掌握号，闭包可以看成函数的高阶应用，是Golang高级开发的必备技能。

匿名函数 匿名函数可以动态的创建，与之成对比的常规函数必须在包中编译前就定义完毕。
匿名函数可以随时改变功能。

“一等公民”意味着函数可以像普通的类型（整型、字符串等）一样进行赋值、作为函数的参数传递、作为函数的返回值等。Golang的函数只能返回匿名函数！

闭包
闭包是匿名函数与匿名函数所引用环境的组合。
匿名函数有动态创建的特性，该特性使得匿名函数不用通过参数传递的方式，就可以直接引用外部的变量。
这就类似于常规函数直接使用全局变量一样，
个人理解为：匿名函数和它引用的变量以及环境，类似常规函数引用全局变量处于一个包的环境。

*/

var f = func(int) {}

func main() {
	f = func(i int) {
		fmt.Println(i)
	}
	f(2)
	f = func(i int) {
		fmt.Println(i * i * i)
	}
	f(2)
	/*
	   输出:
	   2
	   8
	*/

	n := 0
	f := func() int {
		n += 1
		return n
	}
	fmt.Println(f()) // 别忘记括号，不加括号相当于地址
	fmt.Println(f())
	/*
	   输出：
	   1
	   2
	*/

	in := Increase()
	fmt.Println(in())
	fmt.Println(in())
	/*
		输出：
		1
		2
	*/

}

func Increase() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}
