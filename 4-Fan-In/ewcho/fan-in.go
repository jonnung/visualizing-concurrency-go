/*
    https://blog.codeship.com/visualizing-concurrency-go/

	Chapture4. Fan In

    2개의 고루틴은 일정 시간 간격으로 ch 채널로 무한으로 데이터를 송신한다.
	메인 루틴에서는 ch 채널을 range 를 통해 대기하며, 수신된 데이터가 있을시 즉시 out 채널로 다시 송신한다.
	reader 함수가 실행된 고루틴은 out 채널로부터 데이터가 수신되면 결과를 출력한다.
*/
package main

import (
	"fmt"
	"time"
)

func producer(ch chan int, d time.Duration) {
	var i int
	for {
		ch <- i
		i++
		time.Sleep(d)
	}
}

func reader(out chan int) {
	for x := range out {
		fmt.Println(x)
	}
}

func main() {
	ch := make(chan int)
	out := make(chan int)

	go producer(ch, 100*time.Millisecond)
	go producer(ch, 250*time.Millisecond)

	go reader(out)

	// range 문은 일반적인 데이터 구조도 나열할 수 있지만\
	// 채널의 데이터를 나열할 수도 있다.
	for i := range ch {
		out <- i
	}
}
