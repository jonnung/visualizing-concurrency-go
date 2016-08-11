/*
    https://blog.codeship.com/visualizing-concurrency-go/

	Chapter6. Servers

	특정 시간에 request 가 증가할때 로깅 액션이 병목이 될 수 있다.
	요청을 받아 핸들러를 실행하는 고루틴에서 로거가 대기하고 있는 채널에 데이터를 비동기로 전송한다.
*/
package main

import (
	"fmt"
	"net"
	"time"
)

func handler(c net.Conn, ch chan<- string) {
	ch <- c.RemoteAddr().String()
	c.Write([]byte("ok"))
	c.Close()
}

func server(l net.Listener, ch chan string) {
	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}
		go handler(c, ch)
	}
}

func logger(ch <-chan string) {
	for {
		fmt.Println(<-ch)
	}
}

func main() {
	l, err := net.Listen("tcp", ":5000")
	if err != nil {
		panic(err)
	}

	ch := make(chan string)

	go logger(ch)
	go server(l, ch)

	time.Sleep(30 * time.Second)
}
