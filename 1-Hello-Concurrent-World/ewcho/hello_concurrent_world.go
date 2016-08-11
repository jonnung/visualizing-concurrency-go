package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		fmt.Println("Goroutine 1")
		ch <- 12
	}()

	go func() {
		fmt.Println("Goroutine 2")
		ch <- 24
	}()

	go func() {
		fmt.Println("Goroutine 3")
		ch <- 48
	}()

	result1 := <-ch
	result2 := <-ch
	result3 := <-ch

	time.Sleep(time.Second * 1)

	// 고루틴의 동시성은 고루틴들이 동시에 스타트 되는게 아니라 \
	// 고루틴들이 동시(간대)에 실행될 수 있다는 것을 의미하는 것 같다.
	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
}
