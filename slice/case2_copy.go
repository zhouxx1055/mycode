package main

import "fmt"

// 拷贝 copy

func main() {

	s1 := []string{"A", "B", "C"}
	var s2 []string = make([]string, 4)
	fmt.Println(s1)
	fmt.Println(s2)

	copy(s2, s1)

	fmt.Println(s1)
	fmt.Println(s2)

}
