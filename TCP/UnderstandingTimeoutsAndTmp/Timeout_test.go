package UnderstandingTimeoutsAndTmp

import (
	"net"
	"syscall"
	"testing"
	"time"
)

/*
Since the net.DialTimeout functions does not give control of its net.Dialer
to mock the dialer's output, I am making an implementation that matches
the signature.
*/
func DialTimeout(network, address string, timeout time.Duration,
) (net.Conn, error) {

	/* With our DialTimeout mockup function we override the Control function
	of the net.Dialer to return an DNSError, that means it can't
	find the server.*/
	d := net.Dialer{
		Control: func(_, addr string, _ syscall.RawConn) error {
			return &net.DNSError{ //Domain Name Server Error
				Err:         "connection timed out",
				Name:        addr,
				Server:      "127.0.0.1",
				IsTimeout:   true,
				IsTemporary: true,
			}
		},
		Timeout: timeout,
		/* The DialTimout function has the time-out duration
		that we specify in the net.Dialer. A Dialer contains options to connect
		to an address. */
	}
	return d.Dial(network, address)
	/* The DualTimeout function returns the options to connect that is
	the Dial-context (d). Then connects to the network and address. */
}

func TestDialTimeout(t *testing.T) {

	/* Here we try to connect to the address that doesn't exist, and we exceed
	the time-out-duration of 5 seconds.
	If Fatal returns an error that's nil it means that a connection has been made.
	This means that the connection did not time out.*/

	c, err := DialTimeout("tcp", "10.0.0.1:http", 5*time.Second)
	if err == nil {
		c.Close()
		t.Fatal("connection did not time out")
	}
	nErr, ok := err.(net.Error) // We expect a network error
	if !ok {                    // if this is not true we log it
		t.Fatal(err)
	}
	if !nErr.Timeout() { // I check if the error is a Timeout
		t.Fatal("error is not timeout")
	}
}
