package main

import (
	"fmt"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:5555")
	if err != nil {
		fmt.Printf("Error mottat %s", err)
	}

	_, err = conn.Write([]byte("hei echo"))
	if err != nil {
		fmt.Printf("Error mottat %s", err)
	}

	echoBuff := make([]byte, 1024)
	echoStr := ""

	_, err = conn.Read(echoBuff)
	if err != nil {
		fmt.Printf("Error mottat %s", err)
	}

	echoStr += string(echoBuff)

	fmt.Println(echoStr)

	conn.Close()
}
