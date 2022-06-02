package main

import (
	"fmt"
	"io"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:5555")
	if err != nil {
		fmt.Printf("Error motatt13 %d", err)
	}

	defer listener.Close()

	buf := make([]byte, 1024)
	var reply = ""
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error motatt23 %d", err)
		}
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("Error motatt27 %d", err)
		}

		reply = string(buf[:n])
		fmt.Println(reply + "Dette er meldingen fra Klienten")
		reply = deChifferMe(reply)
		fmt.Println(reply)

		_, err = conn.Write([]byte(reply))
		if err != nil {
			if err != io.EOF {
				fmt.Printf("Error motatt36 %d", err)
			}
			break
		}
	}
}

func deChifferMe(input string) string {

	alphabet := [34]rune{' ', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'æ', 'ø', 'å', '.', '0', '1', '2'}

	str := input
	chars := []rune(str)
	newStr := ""

	for i := 0; i < len(chars); i++ {
		for j := 0; j < len(alphabet); j++ {
			if chars[i] == alphabet[j] {
				if j-4 < 0 { //sjekker om den er uten for index
					var a = j - 4
					a += 34
					newStr += string(alphabet[a])
				}
				if !(j-4 < 0) {
					newStr += string(alphabet[j-4])
				}
			}
		}
	}
	return newStr
}
