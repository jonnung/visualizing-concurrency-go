/*
    https://blog.codeship.com/visualizing-concurrency-go/

	Chapture5. Workers

	Fan-In 의 반대 되는 개념의 패턴으로 Fan-Out 또는 Workers 패턴으로 불린다.
	여러개의 고루틴들은 동시에 실행되며, 하나의 채널을 통해 데이터 수신이 될때까지 대기한다.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(tasksCh <-chan int, wg *sync.WaitGroup) { // tasksCh 으로 전달된 채널은 수신만 가능
	defer wg.Done() // defer(지연)을 통해 worker 함수가 끝나기 전에 무조건 WaitGroup 의 Done 메소드가 호출됨

	for {
		task, ok := <-tasksCh // tasksCh 채널로부터 데이터가 수신될때까지 블록 상태

		if !ok {
			return
		}

		d := time.Duration(task) * time.Second
		time.Sleep(d)
		fmt.Println("Processing task", task)
	}
}

func pool(wg *sync.WaitGroup, workers, tasks int) {
	tasksCh := make(chan int)

	for i := 0; i < workers; i++ {
		go worker(tasksCh, wg)
	}

	for i := 0; i < tasks; i++ {
		tasksCh <- i
	}

	close(tasksCh) // 채널을 종료 시키더라도 채널안에 데이터는 존재
}

func main() {
	// WaitGroup은 고루틴들이 완료될때까지 메인 루틴을 기다리게 할 수 있는 메소드 보유
	var wg sync.WaitGroup

	wg.Add(36) // 고루틴이 완료해야할 작업의 양만큼 Add 메소드에 전달

	go pool(&wg, 36, 50)

	wg.Wait() // 모든 고루틴들이 완료될때까지 블록 된다
}
