package main

import (
	"fmt"
	"io"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:5555")
	if err != nil {
		fmt.Printf("Error motatt %d", err)
	}

	defer listener.Close()

	buf := make([]byte, 1024)
	reply := ""
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error motatt %d", err)
		}
		_, err = conn.Read(buf)
		if err != nil {
			fmt.Printf("Error motatt %d", err)
		}
		fmt.Println(string(buf))

		reply += string(buf)

		_, err = conn.Write([]byte(reply))
		if err != nil {
			if err != io.EOF {
				fmt.Printf("Error motatt %d", err)
			}
			break
		}
	}
}
