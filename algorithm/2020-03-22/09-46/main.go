package main

import (
	"fmt"
)

type S struct {
	m string
}

func f() *S {
	return &S{m: "foo"}
}

func main() {
	p := f()
	fmt.Printf("f value %v", p.m)
}
