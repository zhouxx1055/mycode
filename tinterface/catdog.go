package main

import "fmt"

type Cat struct {
	Name string
}
type Dog struct {
	Name string
}

type Mouse struct {
	Name string
}

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
	c := Cat{}
	d := Dog{}
	fmt.Println("猫--->>", c.Say())
	fmt.Println("狗--->>", d.Say())

	c.Name = "Tom"
	d.Name = "Tony"
	fmt.Println("猫--->>", c.Say())
	fmt.Println("狗--->>", d.Say())

	m := Mouse{}
	fmt.Println("老鼠--->>", m.Say())
	m.Name = "Jerry"
	fmt.Println("老鼠--->>", m.Say())

}
