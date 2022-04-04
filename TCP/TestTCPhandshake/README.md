## Establishing a TCP Connection by Using Go's Standard Library.

* By using the net package in Go's standard library i can create TCP-based servers and clients capble of connecting to those servers.
***
***Binding, Listening for, and Accepting Connections***
* To create a TCP server capable of listening for incoming connections (called a listener), use the net.Listen function.
* This function will retrun an object that implements the net.Listener interface.


***Source***
- Adam Woodbeck - Network Programming with Go_ Learn to Code Secure and Reliable Network Services from Scratch (2021, No Starch Press)  

___
@Server

"The listener accepts a network type ("tcp") and an IP address
and port separated by colon  ("127.0.0.1:0").
The function returns a net.Listener interface and an error interface. If the function
returns successfully, the listener is bound to the specified IP address and port.
Binding means that the operating system has exclusively assigned the port on the given IP address
to the listener. If the IP address is occupied it will return an error. If the param
is empty, Go will assign a random port number to the listener.

You can retrieve the IP address by calling its Addr method, listener.Addr().
Likewise, if you commit the IP address, your listener will be bound to all unicast
and any anycast IP addresses on the system. Omitting both IP address and port, or
passing in a colon for the second argument to net.Listener, will cause you listener to
bind all uncast and anycast IP addresses using a random port.

In most cases, you should use tcp as the network type for net.Listener's first argument.
It's possible to restrict the listener to just IPv4 addresses by passing tcp4 or IPv6 addresses
by passing tcp6.

You should be diligent about closing you listener gracefully by calling its Close method,
often in a defer if it makes sense for your code. Failure to close the listener
may lead to memory leaks or deadlocks in your code, because calls to the listener's Accept
method may block indefinitly. Closing the listener immediately unblocks calls to the Accept method."

Unless you want to accept only a single incoming connection, you need to use a for loop
so your server will accept each incoming connection, handle it in a goroutine, and loop back around,
ready to accept the next connection.

We start the for loop by calling the listener's incoming connection and completes the TCP handshake process.
This method will block until the listener detects an incoming connection and completes the TCP handshake process
between  the client and the server. The call returns a net.Conn interface and an error.

The connection interface's is a pointer to a net.TCPConn object because you're accepting TCP connections.
The connection interface represents the server's side of the TCP connection. In most cases, net.Conn provides all
methods you'll need for general interactions with the client.

The for loop lets us concurrently handle client connections. You spin of a goroutine to
asynchronously handle each connection so your listener can ready itself for the next client. Then we call
the connection's Close method before the goroutine exits to terminate the connections by sending a FIN packet to the server.

***The goroutine (go socket)***
```go
for	{
conn, err := listener.Accept()
if err != nil {
return err
}
go func (c net.Conn) {
defer c.Close()

	}(conn)
}
```
@Client

Start by creating a listener on the IP address 127.0.0.1, which the client will connect to.
Then spin off the listener in a goroutine so i can work with the client's side of the connection later in the test.

The standard library's net.Dial function is like the net.Listen function in that it accepts a network like tcp
and an IP address and port combination to witch it's trying to connect.

You can use a hostname in place of an IP address and a service name, like http, in place of a port number.
If a hostname resolves to more than one IP address, Go will attempt a connection to each one in order until
a connection succeeds or all IP addresses have been exhausted.

After the successful connection to the listener, I initiate a termination of the connection from the
client's side. After receving the finish (FIN) packet, the Read method returns the io.EOF error,
indicating to the listener's code that you closed your side of the connection. The connection
handler exits, calling the connection's Close method. This sends a Fin packet to your connection, completing
the termination of the TCP session.

Then I close the listener. The listener's Accept method immediately unblocks and returns an error. This error
isn't a failure, because we expect it, so I log it using t.Logf. The listener's goroutine
will then exit.
