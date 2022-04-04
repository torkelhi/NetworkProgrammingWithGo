package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:6000")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()
	buf := make([]byte, 1)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			break
		}

		if n > 0 {
			fmt.Printf("read %d bytes\n", n)
			fmt.Println(n)
			fmt.Println(err)
		}
	}

	/*
		Read - interfacen kan brukes til mye forskjellig.
		Her brukes den for å lese gjennom README.md filen"
	*/
	f, err := os.Open("ReadInterface_test.md")
	if err != nil {
		panic(err)
	}
	file := make([]byte, 5)
	for {
		read, err := f.Read(file)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf(string(file[:read]))
	}
}
