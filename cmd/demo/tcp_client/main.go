package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9090")
	defer conn.Close()
	if err != nil {
		fmt.Printf("connect failed: %v\n", err)
		return
	}

	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("read from console failed: %v\n", err)
			break
		}

		trimmedInput := strings.TrimSpace(input)
		if trimmedInput == "Q" {
			break
		}

		_, err = conn.Write([]byte(trimmedInput))
		if err != nil {
			fmt.Printf("write failed: %v\n", err)
			break
		}

		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("read from connect failed, err: %v\n", err)
			break
		}

		str := string(buf[:n])
		fmt.Printf("received %v\n", str)

	}

}
