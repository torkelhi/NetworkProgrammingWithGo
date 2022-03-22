package TestTCPhandshake

import (
	"io" //stands for I/O- input and output module.
	"net"
	"testing"
)

func TestDial(t *testing.T) {
	// creating a listener on a random port, empty param with colon (:)

	listener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}

	done := make(chan struct{})
	go func() {
		defer func() { done <- struct{}{} }() //defer: does the defer function last.

		for {
			conn, err := listener.Accept()
			if err != nil {
				t.Log(err)
				return
			}
			// go func() - Anomaly function, can be used as param/variable. Is called at through ().
			go func(c net.Conn) {
				defer func() {
					c.Close()
					done <- struct{}{}
				}() // go func (c net.Conn)

				buf := make([]byte, 1024)
				for {
					n, err := c.Read(buf) //Read() reads the data for the connection.
					if err != nil {
						if err != io.EOF { //EOF is the error returned by Read when no more input is available.
							t.Error(err)
						}
						return
					}
					//Logf formats its arguments according to the format, analogous to Printf, and records the text in the error log.
					//A final newline is added if not provided.
					//For tests, the text will be printed only if the test fails or the -test.v flag is set
					t.Logf("received: %q", buf[:n])
				}
			}(conn)
		}
	}()

	//Dial() connects to the address on the named network.
	//Addr() returns the listeners network and String() turns it to a String datatype.
	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}

	//Closes the connection between server and client. Sends finsh (FIN).
	conn.Close()
	<-done
	listener.Close()
	<-done
}
