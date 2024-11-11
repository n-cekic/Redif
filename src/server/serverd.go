package main

import (
	"fmt"
	"net"
	L "redif/common/logger"
)

func main() {
	L.Init()

	port := 8080
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		L.Fatal("failed listening on port ", port, ". Error: ", err.Error())
		return
	}

	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			L.Fatal("failed listening on port ", port, ". Error: ", err.Error())
			return
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

}
