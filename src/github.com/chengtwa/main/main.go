package main

import (
	"fmt"
)

func main() {
	// var v1 bool
	// v1 = true
	// v2 := (1 == 2)
	// fmt.Println(v1, v2, v1)
	var msg = "aaa"
	foo := "bbb"
	foo1 := []rune(foo)
	foo1[0] = 'c'
	// bar := func() {
	// 	foo = "ccc"
	// }
	// bar()
	fmt.Println("string", msg)
	fmt.Println("foo:", foo)
	fmt.Println("foo1:", string(foo1))
}
