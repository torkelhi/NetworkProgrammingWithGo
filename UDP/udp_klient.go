package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	serverAddr, err := net.ResolveUDPAddr("udp", "localhost:5656")
	if err != nil {
		log.Fatalf("clientAddr --> Resieved error: %b", err)
	}

	localAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:")
	if err != nil {
		log.Fatalf("localAddr --> Resieved error: %b", err)
	}
	//localAddr := nil

	//localAddr kan være nil fordi DialUDP() antar at man skal bruke en localAddr.
	conn, err := net.DialUDP("udp", localAddr, serverAddr)
	if err != nil {
		log.Fatalf("Resieved error: %s", err)
	}

	defer conn.Close()

	//denne meldingen skal bli sendt tilbake fra server.
	clientMsg := []byte("Greetings World! æøå")

	_, err = conn.Write(clientMsg)
	//Write() returnerer int og error. Trenger ikke int --> _, err.
	if err != nil {
		log.Printf("WriteToUDP --> Recieved: %s", err)
	}

	//lager buffer for å motta clientMsg tilbake fra server.
	echoBuff := make([]byte, 1024)

	n, addr, err := conn.ReadFromUDP(echoBuff)
	if err != nil {
		log.Printf("Client Read--> Recieved: %s", err)
	}

	fmt.Println("Client Recieved: ", addr)         //Server sin socket
	fmt.Println("Message: ", string(echoBuff[:n])) //echoMsg
	fmt.Println("Number of bytes: ", n)            //number of bytes.

}
