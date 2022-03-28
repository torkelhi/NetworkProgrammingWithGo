package TestTCPhandshake

import (
	"crypto/rand"
	"io"
	"net"
	"testing"
)

/*
Test for å sende mock data ved hjelp av crypto/rand package fra klient
til en server med 3-way-handshake.
*/

func TestReadIntoBuffer(t *testing.T) {
	// genererer mock data

	payload := make([]byte, 1<<24) //allokerer ca. 16777216 bytes.

	_, err := rand.Read(payload) // legger inn tilfeldige byte (rand), leser de av (Read).
	if err != nil {
		t.Fatal(err)
	}

	listener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		t.Fatal(err) // Samme som t.Log(), men følger opp med en "end now".
	}

	go func() {
		conn, err := listener.Accept()
		if err != nil {
			t.Log(err) // viser hvilke error som har hendt.
			return
		}
		defer conn.Close()

		_, err = conn.Write(payload) // Write() skriver data til tilkoblingen
		if err != nil {
			t.Error(err) // samme som t.Log(), men markerer testen med at den failet her.
		}
	}()

	conn, err := net.Dial("tcp", listener.Addr().String())

	buf := make([]byte, 1<<19) // 512 Kilobytes

	for {
		n, err := conn.Read(buf)

		if err != nil {
			if err != io.EOF {
				t.Error(err)
			}
			break
		}
		t.Logf("read %q bytes", n)
	}
}
