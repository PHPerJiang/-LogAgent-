package main

import (
	"fmt"
)

//回文
func handle(content string) bool {
	if len(content) == 0 {
		return false
	}
	content_rune := []rune(content)
	i, j := 0, len(content_rune)-1
	for i < j {
		if content_rune[i] == content_rune[j] {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

func main() {
	content := `上海自来水来自海上`
	fmt.Printf("%v\n", handle(content))
}
