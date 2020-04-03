package main

import (
	"fmt"
	"sync"
)

//使用两个协程交替输出20以内的奇偶数，不允许使用临时变量

var (
	chan1 = make(chan int)
	chan2 = make(chan int)
	wg    = &sync.WaitGroup{}
)

func goroutine1(wg *sync.WaitGroup, chan1, chan2 chan int) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		chan2 <- 2*i - 1
		fmt.Printf("this is goroutine 1 : %#v\n", <-chan1)
	}
}

func goroutine2(wg *sync.WaitGroup, chan1, chan2 chan int) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		fmt.Printf("this is goroutine 2 : %#v\n", <-chan2)
		chan1 <- 2 * i
	}
}

func main() {
	wg.Add(2)
	go goroutine1(wg, chan1, chan2)
	go goroutine2(wg, chan1, chan2)
	wg.Wait()
}
