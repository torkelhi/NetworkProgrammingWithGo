package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:5099")
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		defer conn.Close()

		var messagefromserverBuffer []string
		scanner := bufio.NewScanner(conn)
		scanner.Split(bufio.ScanWords)

		for scanner.Scan() {
			messagefromserverBuffer = append(messagefromserverBuffer, scanner.Text())
		}

		fmt.Printf("Scanned words: %#v\n", messagefromserverBuffer)

		_, err := conn.Write([]byte("halo, client speaking her"))
		if err != nil {
			fmt.Println(err)
			if err != io.EOF {
				fmt.Println(err)
			}
		}
	}()
}
