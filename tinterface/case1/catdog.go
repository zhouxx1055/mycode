package main

import "fmt"

/**
接口（interface）定义了一个对象的行为规范，只定义规范不实现，由具体的对象来实现规范的细节。
*/
// 接口定义
type Sayer interface {
	Say() string
}

// 对象定义 （go中：接口是种类型，一种抽象的类型。）
type Cat struct {
	Name string
	Sayer
}
type Dog struct {
	Name string
	Sayer
}

type Mouse struct {
	Name string
	Sayer
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

func main() {
	var x Sayer                    // 声明一个Sayer类型的变量x
	a := Cat{}                     // 实例化一个cat
	b := Dog{}                     // 实例化一个dog
	x = a                          // 可以把cat实例直接赋值给x
	fmt.Println("猫--->>", x.Say()) // 喵喵喵
	x = b                          // 可以把dog实例直接赋值给x
	fmt.Println("狗--->>", x.Say()) // 汪汪汪

	x = a // 可以把cat实例直接赋值给x
	a.Name = "Tom"
	fmt.Println("猫--->>", x.Say()) // 喵喵喵
	x = b                          // 可以把dog实例直接赋值给x
	b.Name = "Tony"
	fmt.Println("狗--->>", x.Say()) // 汪汪汪

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
