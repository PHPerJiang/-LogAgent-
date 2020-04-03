package main

import (
	"fmt"
	"time"
)

/**
* 从1000里找到  a*a + b*b = c*c 的所有数，并且不允许使用函数包及math包
* 要求性能要高
* baseVersion           耗时:119647000
* optimizationVersion1  耗时:1401000
* optimizationVersion2  耗时:508000
 */

func baseVersion() {
	start := time.Now().Nanosecond()
	for a := 0; a <= 1000; a++ {
		for b := 0; b <= 1000; b++ {
			for c := 0; c <= 1000; c++ {
				if (a+b+c) == 1000 && (a*a+b*b) == c*c {
					fmt.Printf("%d  %d  %d\n", a, b, c)
				}
			}
		}
	}
	fmt.Printf("耗时:%v\n", time.Now().Nanosecond()-start)
}

func optimizationVersion1() {
	start := time.Now().Nanosecond()
	for a := 0; a <= 1000; a++ {
		for b := 0; b <= 1000; b++ {
			c := 1000 - a - b
			if a*a+b*b == c*c {
				fmt.Printf("%d  %d  %d\n", a, b, c)
			}
		}
	}
	fmt.Printf("耗时:%v\n", time.Now().Nanosecond()-start)
}

func optimizationVersion2() {
	start := time.Now().Nanosecond()
	for a := 0; a <= 1000; a++ {
		for b := 0; b <= 1000-a; b++ {
			c := 1000 - a - b
			if a*a+b*b == c*c {
				fmt.Printf("%d  %d  %d\n", a, b, c)
			}
		}
	}
	fmt.Printf("耗时:%v\n", time.Now().Nanosecond()-start)
}

func main() {
	baseVersion()
	optimizationVersion1()
	optimizationVersion2()
}
