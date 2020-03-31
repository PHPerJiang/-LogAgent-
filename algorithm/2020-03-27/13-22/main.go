package main

import (
	"fmt"
)

func main() {
	var a = []string{"a", "b", "c"}
	for v := range a {
		fmt.Print(v)
	}
}
