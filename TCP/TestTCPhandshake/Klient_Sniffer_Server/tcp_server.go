package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:6669") //listener on port 6669
	if err != nil {
		fmt.Println(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go func() {
			_, err = conn.Write([]byte("Welcome to my server"))
			if err != nil {
				fmt.Println(err)
			}
		}()
	}
}
