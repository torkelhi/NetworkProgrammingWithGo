package main

import (
	"crypto/rand"
	"fmt"
	"net"
)

func main() {

	// Genererer mock data
	payload := make([]byte, 20) // alloker 16777216 bytes

	_, err := rand.Read(payload) // leser inn tilfeldige byte i payload
	if err != nil {
		fmt.Println(err)
	}

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
		_, err = conn.Write(payload)
		if err != nil {
			fmt.Println(err)
			break
		}
		conn.Close()
	}
}
