package main

import "fmt"

/*
切片不保存具体值
切片对应一个底层数组
底层数组都是占用一块连续的内存
*/
func main() {
	s1 := []string{"A", "B", "C", "D", "E"}
	fmt.Println(s1)
	fmt.Println(len(s1))
	s1 = append(s1[:1], s1[2:]...) // 删除index=1
	fmt.Println(s1)
	fmt.Println(len(s1))

	fmt.Println(s1)
	fmt.Println(len(s1))
	s1 = delete(s1, 1) // 删除index=1
	fmt.Println(s1)
	fmt.Println(len(s1))

}

func delete(s1 []string, index int) []string {
	if index < 0 || index >= len(s1) {
		return s1
	}
	return append(s1[:index], s1[index+1:]...)
}
