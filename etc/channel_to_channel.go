package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	out := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			i++
			time.Sleep(time.Second * 1)
		}
	}()

	out <- ch

	for data := range out {
		fmt.Println(data)
	}
}
