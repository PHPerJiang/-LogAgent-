package main

import (
	"fmt"
	"log"
)

const N = 3

func main() {
	m := make(map[int]*int)
	for i := 0; i < N; i++ {
		log.Println(N, i, &i)
		m[i] = &i
	}
	for _, v := range m {
		fmt.Println(*v)
	}
}
