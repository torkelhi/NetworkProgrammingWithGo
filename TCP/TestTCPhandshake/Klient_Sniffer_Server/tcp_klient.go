package main

import (
	"net"
	"os"
)

func main() {
	strEcho := ""
	servAddr := "127.0.0.1:6669"

	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr) //net.ResolveTCPAddr returnerer TCP adressen, når man gir server adressen
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr) //"Ringer etter en TCP server
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}

	reply := make([]byte, 1024) //allokerer plass i minne for data fra server

	_, err = conn.Read(reply) //Leser av data sendt fra server
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}
	println("reply from server=", string(reply))
	strEcho += string(reply)

	_, err = conn.Write([]byte(strEcho)) //skriver til koblingen vi etablerte
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	println("write to server = ", strEcho)

	conn.Close() //lukker tilkoblingen
}
