package main

import "log"

func main() {
	defer func() {
		log.Println("defer 1")
	}()

	defer func() {
		if err := recover(); err != nil {
			log.Printf("catch panic,%v", err)
		}
	}()

	panic("this is panic")
}
