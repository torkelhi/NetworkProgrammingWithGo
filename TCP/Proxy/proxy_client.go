package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	//1. Dial proxy_proxy_server - net.Dial
	//2. Get redirected to proxy_server
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print(conn.LocalAddr()) //checks the address (should be port 6666)

	buf := make([]byte, 10)
	defer conn.Close()

	//3. Read incoming data from conn (proxy_server) - read interface

	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			break
		}

		if n < 0 {
			fmt.Printf(string(buf[:n]))
		}
	}
	//4. Close connection. (gracefull close) - conn.Close | listener.Close
}
