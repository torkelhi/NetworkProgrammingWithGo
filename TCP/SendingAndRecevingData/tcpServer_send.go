package main

import (
	"fmt"
	"io"
	"net"
)

func main() {

	const payload = "Halo, server speaking here"

	listener, err := net.Listen("tcp", "127.0.0.1:5099")
	if err != nil {
		fmt.Println(err)
	}

	for {
		var buf = make([]byte, 1<<16)
		conn, err := listener.Accept()
		go func() {
			if err != nil {
				fmt.Println(err)
				return
			}
			_, err = conn.Write([]byte(payload))
			if err != nil {
				if err != io.EOF {
					fmt.Println(err)
				}
			}
			conn.Read(buf)
			conn.Close()
		}()
	}
}
