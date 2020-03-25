package main

import "log"

var a interface{}

func main() {
	a = int(1)
	value, ok := a.(int)
	if !ok {
		log.Printf("isnt string,err:%v", ok)
	}
	log.Print(value)
}
