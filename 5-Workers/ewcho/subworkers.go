package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

const (
	WORKERS    = 4
	SUBWORKERS = 2
	TASKS      = 10000
	SUBTASKS   = 100
)

func worker(tasks chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		task, ok := <-tasks
		if !ok {
			return
		}

		subtasks := make(chan int)

		for i := 0; i < SUBWORKERS; i++ {
			go subworker(subtasks)
		}

		for i := 0; i < SUBTASKS; i++ {
			subtask := task * i
			subtasks <- subtask
		}
		close(subtasks)
	}
}

func subworker(subtask chan int) {
	for {
		task, ok := <-subtask
		if !ok {
			return
		}
		time.Sleep(time.Duration(task) * time.Millisecond)
		fmt.Println(task)
	}
}

func main() {
	// Go 는 기본적으로 CPU 하나만 사용하게 설계됨
	// GOMAXPROCS 함수를 이용해서 런타임상에 사용할 CPU 개수를 지정할 수 있음
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Printf("Usable CPU Core: %d", runtime.GOMAXPROCS(0))
	time.Sleep(5 * time.Second)

	var wg sync.WaitGroup

	wg.Add(WORKERS)

	tasks := make(chan int)

	for i := 0; i < WORKERS; i++ {
		go worker(tasks, &wg)
	}

	for i := 0; i < TASKS; i++ {
		tasks <- i
	}

	close(tasks)
	wg.Wait()
}
