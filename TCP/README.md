##Establishing a TCP Connection by Using Go's Standard Library.

* By using the net package in Go's standard library i can create TCP-based servers and clients capble of connecting to those servers.
***
***Binding, Listening for, and Accepting Connections***
* To create a TCP server capable of listening for incoming connections (called a listener), use the net.Listen function.
* This function will retrun an object that implements the net.Listener interface.


***Source***
- Adam Woodbeck - Network Programming with Go_ Learn to Code Secure and Reliable Network Services from Scratch (2021, No Starch Press)