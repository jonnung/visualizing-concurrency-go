/*
    https://blog.codeship.com/visualizing-concurrency-go/

	Chapter6. Servers

	Listener를 생성하고, 루프에서 Accept() 함수를 실행하고 커넥션이 들어오면
	핸들러 함수를 고루틴으로 커넥션을 처리한다.
*/
package main

import "net"

func handler(c net.Conn) {
	c.Write([]byte("ok"))
	c.Close()
}

func main() {
	l, err := net.Listen("tcp", ":5000")
	if err != nil {
		panic(err)
	}

	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}
		go handler(c)
	}
}
