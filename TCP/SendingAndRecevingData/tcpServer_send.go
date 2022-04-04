package main

import (
	"fmt"
	"net"
)

func main() {

	payload := "halo, server speaking here."

	listener, err := net.Listen("tcp", "127.0.0.1:6000")
	if err != nil {
		fmt.Println(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			break
		}
		_, err = conn.Write([]byte(payload))
		if err != nil {
			fmt.Println(err)
			break
		}
		conn.Close()
	}
}
