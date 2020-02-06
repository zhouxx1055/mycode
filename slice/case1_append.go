package main

import "fmt"

// 越界 ,append

func main() {
	s1 := []string{"A", "B", "C"}
	fmt.Println(s1)
	fmt.Println(len(s1))
	//s1[3]="D" //编译不报错  运行报错  panic: runtime error: index out of range
	s1 = append(s1, "D")
	fmt.Println(s1)
	fmt.Println(len(s1))

	var s2 = make([]string, 4)
	s2[0] = "A1"
	s2[1] = "B1"
	s2[2] = "C1"
	fmt.Println(s2)
	fmt.Println(len(s2)) // 长度4
	s2 = append(s2, "D1")
	fmt.Println(s2)
	fmt.Println(len(s2)) // 长度5

	s1 = append(s1, s2...)
	fmt.Println(s1)
	fmt.Println(len(s1))
}
