package main

import (
	"fmt"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:5555")
	if err != nil {
		fmt.Printf("Error mottat12 %s", err)
	}

	_, err = conn.Write([]byte(chifferMe("w, x og y møtes i ålesund")))
	if err != nil {
		fmt.Printf("Error mottat17 %s", err)
	}

	echoBuff := make([]byte, 1024)
	echoStr := ""

	n, err := conn.Read(echoBuff)
	if err != nil {
		fmt.Printf("Error mottat25 %s", err)
	}
	fmt.Println(echoBuff[:n])
	echoStr += string(echoBuff[:n])

	fmt.Println(echoStr)

	conn.Close()
}

func chifferMe(input string) string {

	alphabet := [31]rune{' ', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'æ', 'ø', 'å', ','}

	str := input
	chars := []rune(str)
	newStr := ""

	for i := 0; i < len(chars); i++ {
		for j := 0; j < len(alphabet); j++ {

			if chars[i] == alphabet[j] {
				if j+4 > len(alphabet) {
					var a = j + 4
					a -= 31
					newStr += string(alphabet[a])
				}
				if !(j+4 > len(alphabet)) {
					newStr += string(alphabet[j+4])
				}
			}
		}
	}
	return newStr
}
