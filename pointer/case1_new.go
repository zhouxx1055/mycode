package main

import "fmt"

// new 给指针分配内存
/**
// The new built-in function allocates memory. The first argument is a type,
// not a value, and the value returned is a pointer to a newly
// allocated zero value of that type.
func new(Type) *Type
*/
func main() {
	n := 20
	fmt.Println(&n)
	fmt.Println(*(&n))

	//var a *int
	//*a = 20 // 报错 panic: runtime error: invalid memory address or nil pointer dereference
	//fmt.Println(*a)

	var a = new(int)
	*a = 20
	fmt.Println(*a)
	fmt.Println(&a)
	fmt.Println(a)

	//var b =new(int)
	//b = 20 // 报错  cannot use 20 (type int) as type *int in assignment
	//fmt.Println(*b)
	//fmt.Println(&b)
	//fmt.Println(b)
}
