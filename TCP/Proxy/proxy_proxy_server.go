package main

import (
	"io"
	"log"
	"net"
)

func main() {
	//1. listen for proxy request
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatalln("Unable to bind on port")
		// Fatalln is equivalent to Println() followed by a call to os.Exit(1)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}

		//2. Handle the proxy request
		go handleConn(conn)
	}
}

func handleConn(src net.Conn) {
	//1. Connect to the target server
	dst, err := net.Dial("tcp", "127.0.0.1:6666")
	if err != nil {
		log.Fatalln(err)
	}

	defer dst.Close()

	//2. Copy the src output to the destination conn (dst)

	go func() {
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()

	//3. Copy the respons from target for server conn to src conn

	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}
}
