package main

import "fmt"

/**
  1. 取值 *
  2. 取地址 &
*/
func main() {
	n := 20
	fmt.Println(&n)
	fmt.Println(*(&n))

}
