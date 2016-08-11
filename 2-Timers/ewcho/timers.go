package main

import (
	"fmt"
	"time"
)

func timer(d time.Duration) <-chan int {
	c := make(chan int)

	go func() {
		time.Sleep(d)
		c <- 1
	}()

	return c
}

func main() {
	for i := 0; i < 24; i++ {
		fmt.Println()
		fmt.Printf("Goroutine start: %d", i)
		fmt.Println()

		c := timer(1 * time.Second)

		result := <-c // main() 루틴에서 c 채널에서 데이터를 계속 수신을 대기하고 있음

		fmt.Println(result)
		fmt.Printf("Goroutine end: %d", i)
		fmt.Println()
	}
}
