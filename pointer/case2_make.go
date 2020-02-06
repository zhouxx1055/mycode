package main

import "fmt"

/** 、、 给slice, map, or chan (only)分配内存
// The make built-in function allocates and initializes an object of type
// slice, map, or chan (only). Like new, the first argument is a type, not a
// value. Unlike new, make's return type is the same as the type of its
// argument, not a pointer to it. The specification of the result depends on
// the type:
//	Slice: The size specifies the length. The capacity of the slice is
//	equal to its length. A second integer argument may be provided to
//	specify a different capacity; it must be no smaller than the
//	length. For example, make([]int, 0, 10) allocates an underlying array
//	of size 10 and returns a slice of length 0 and capacity 10 that is
//	backed by this underlying array.
//	Map: An empty map is allocated with enough space to hold the
//	specified number of elements. The size may be omitted, in which case
//	a small starting size is allocated.
//	Channel: The channel's buffer is initialized with the specified
//	buffer capacity. If zero, or the size is omitted, the channel is
//	unbuffered.
func make(t Type, size ...IntegerType) Type
*/
func main() {
	n := 20
	fmt.Println(&n)
	fmt.Println(*(&n))

	var map1 map[string]int
	//map1["张三"]=26 //报错 panic: assignment to entry in nil map
	fmt.Println(map1)

	map1 = make(map[string]int, 3)
	map1["张三"] = 26

	fmt.Println(map1)

	var strArray []string
	//strArray[0]="A" //panic: runtime error: index out of range
	fmt.Println(strArray)
	strArray = make([]string, 3)
	strArray[0] = "A"
	fmt.Println(strArray)
}
