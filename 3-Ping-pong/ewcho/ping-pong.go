/*
	https://blog.codeship.com/visualizing-concurrency-go/

	Chapture3. Ping-Pong

	숫자를 가지고 핑퐁을 한다.
	숫자는 1개의 채널을 통해서 송/수신 된다.
	2개의 고루틴 플레이어가 존재한다.
	메인 루틴에서 고루틴을 실행하고,
	2개의 고루틴에서 먼저 채널에서 값을 받아온 쪽이 값을 증가 시키고 다시 채널로 보낸다.
*/
package main

import (
	"fmt"
	"time"
)

func player(ch chan int) {
	for {
		ball := <-ch // 채널로부터 데이터가 수신될때까지 대기
		ball++
		ch <- ball // 다시 채널로 데이터 전송

		if ball%2 != 0 {
			fmt.Println("-------> o")
		} else {
			fmt.Println("o <-------")
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	fmt.Println("Reday...")
	ch := make(chan int)

	go player(ch) // 고루틴 1
	go player(ch) // 고루틴 2

	fmt.Println("Go!!")
	ch <- 0

	// 1초 뒤에 채널에서 데이터를 수신 하면서 위 고루틴 1, 2는 다시 대기하게 된다
	time.Sleep(1 * time.Second)
	result := <-ch

	fmt.Println("======== END ========")
	fmt.Printf("Rally count: %d", result)
}
