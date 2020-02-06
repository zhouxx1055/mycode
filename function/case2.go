package main

import "fmt"

// 高阶函数   ：函数作为参数或者返回值

type student struct {
	name  string
	grade int8
}

func filter(stu []student, f func(s student) bool) []student {
	var r []student

	for _, s := range stu {
		if f(s) == true {
			r = append(r, s)
		}
	}

	return r
}

func main() {
	s1 := student{
		"zhangsan",
		90,
	}

	s2 := student{
		"lisi",
		80,
	}

	s3 := student{
		"wanggang",
		70,
	}

	s := []student{s1, s2, s3}

	fmt.Println("all student: ", s)

	var result []student

	result = filter(s, func(s student) bool {
		if s.grade < 90 {
			return true
		}
		return false
	})

	fmt.Println("less than 90: ", result)
}
