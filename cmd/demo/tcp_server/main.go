package main

import (
	"fmt"
	"net"
	"net/http"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		fmt.Printf("listen fail, err: %v\n", err)
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("accept fail, err: %v\n", err)
			continue
		}

		go process(conn)
	}
}

func process(conn net.Conn) {
	http.NewServeMux()
	defer conn.Close()
	for {
		var buf [8]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("read from connect failed, err: %v\n", err)
			break
		}

		str := string(buf[:n])
		fmt.Printf("received %v\n", str)

		conn.Write([]byte("=============== ok ===============\n"))
	}
}
