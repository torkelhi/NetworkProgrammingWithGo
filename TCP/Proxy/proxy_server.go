package main

import (
	"log"
	"net"
)

func main() {
	//1. Listen for incoming request - Listener

	listener, err := net.Listen("tcp", "127.0.0.1:6666")
	if err != nil {
		log.Fatalln(err)
	}

	payload := "proxy_server speaking here"
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		//2. Send data - Write interface
		_, err = conn.Write([]byte(payload))
		if err != nil {
			log.Fatalln()
		}
		conn.Close()
	}
}
