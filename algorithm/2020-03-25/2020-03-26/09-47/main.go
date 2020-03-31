package main

import "fmt"

var i int = 1
var m = 1

func main() {
	fmt.Printf("%T\n", i)
	fmt.Printf("%T\n", m) //m为自动推断类型
}
