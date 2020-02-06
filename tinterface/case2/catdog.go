package main

import "fmt"

// 接口定义
type Sayer interface {
	Say() string
}

// 抓老鼠
type CatchMiceer interface {
	CatchMice()
}

// 看家
type Watchdoger interface {
	Watchdog()
}

// 对象定义 （go中：接口是种类型，一种抽象的类型。）
type Animal struct {
	Name string
	Sayer
}
type Cat struct {
	Animal
	CatchMiceer
}
type Dog struct {
	Animal
	Watchdoger
}

type Mouse struct {
	Animal
}

// 接口实现
func (c Cat) Say() string {
	if c.Name == "" {
		return "喵。。。"
	}
	return "我是" + c.Name + ",喵。。。"
}
func (d Dog) Say() string {
	if d.Name == "" {
		return "旺。。。"
	}
	return "我是" + d.Name + ",旺。。。"
}
func (d Mouse) Say() string {
	if d.Name == "" {
		return "叽。。。"
	}
	return "我是" + d.Name + ",叽。。。"
}

func (c Cat) CatchMice() {
	if c.Name == "" {
		fmt.Println("我可以抓老鼠！")
	} else {
		fmt.Println("我是" + c.Name + ",我可以抓老鼠！")
	}
}
func (d Dog) Watchdog() {
	if d.Name == "" {
		fmt.Println("我可以看家！")
	} else {
		fmt.Println("我是" + d.Name + ",我可以看家！")
	}
}

func main() {
	var x Sayer                    // 声明一个Sayer类型的变量x
	a := Cat{}                     // 实例化一个cat
	b := Dog{}                     // 实例化一个dog
	x = a                          // 可以把cat实例直接赋值给x
	fmt.Println("猫--->>", x.Say()) // 喵喵喵
	a.CatchMice()
	x = b                          // 可以把dog实例直接赋值给x
	fmt.Println("狗--->>", x.Say()) // 汪汪汪
	b.Watchdog()

	x = a // 可以把cat实例直接赋值给x
	a.Name = "Tom"
	fmt.Println("猫--->>", x.Say()) // 喵喵喵
	a.CatchMice()
	x = b // 可以把dog实例直接赋值给x
	b.Name = "Tony"
	fmt.Println("狗--->>", x.Say()) // 汪汪汪
	b.Watchdog()

	a.Name = "Tom"
	x = a                          // 可以把cat实例直接赋值给x
	fmt.Println("猫--->>", x.Say()) // 喵喵喵
	b.Name = "Tony"
	x = b                          // 可以把dog实例直接赋值给x
	fmt.Println("狗--->>", x.Say()) // 汪汪汪

	x = &a // 可以把cat实例地址 赋值给x
	a.Name = "Tom"
	fmt.Println("猫--->>", x.Say()) // 喵喵喵
	x = &b                         // 可以把dog实例地址 赋值给x
	b.Name = "Tony"
	fmt.Println("狗--->>", x.Say()) // 汪汪汪
}
