package main

import (
	"fmt"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:6000")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()
	buf := make([]byte, 1)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			break
		}

		if n > 0 {
			fmt.Printf(string(buf[:n]))
		}
	}
}
