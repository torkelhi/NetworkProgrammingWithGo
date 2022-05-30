package main

import (
	"log"
	"net"
)

func udpServerHandler(conn *net.UDPConn) {

	buff := make([]byte, 1024)
	echoStr := ""

	n, clientAddr, err := conn.ReadFromUDP(buff)
	if err != nil {
		log.Printf("Reader recieved error: %b", err)
	}

	log.Println("UDP_client: ", clientAddr)                       //Client sin socket. IP = 127.0.0.1 : Port: ulik for hver klient.
	log.Println("Received from UDP client :  ", string(buff[:n])) //Melding fra Klient
	log.Println("Number of bytes: ", n)

	//handler tar seg av kommunikasjonen.
	//sender tilbake stringen mottatt av klient
	echoStr += string(buff[:n]) //

	msg := []byte(echoStr)
	_, err = conn.WriteToUDP(msg, clientAddr)

	if err != nil {
		log.Println(err)
	}
}

func main() {

	serverAddr, err := net.ResolveUDPAddr("udp", "localhost:5656")
	if err != nil {
		log.Fatalf("Error recieved error: %s", err)
	}

	listener, err := net.ListenUDP("udp", serverAddr)
	if err != nil {
		log.Fatalf("Error recieved error: %s", err)
	}
	defer listener.Close()

	//forbindelsesløs --> trenger ingen logisk forbindelse. ingen Accept().
	for {
		udpServerHandler(listener)
	}
}
