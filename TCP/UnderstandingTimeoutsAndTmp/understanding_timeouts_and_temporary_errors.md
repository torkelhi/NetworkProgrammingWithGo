#Time-outs and tmp-errors

***Source***
- Adam Woodbeck - Network Programming with Go_ Learn to Code Secure and Reliable Network Services from Scratch (2021, No Starch Press)
___
Connections will not always succeed. Therefor we determine if an error is tmp or  
if it should be terminated. Go's net package provides interfaces to get more insight  
on the specific error. The net.Error interface includes two important methods.  
Timeouts and Temporary.  

***The Timeout method***  
The Timeout method returns true on Unix-based OP and Windows  
if the OP tells Go that the resource us temporarily unavailable, the call would block,  
or the connection timed out.  

***The Temporary method***  
Returns true if the error's Timeout function returns true, the function call was interrupted,  
or there are too many open files on the system, usually because OP has exceeded  
the system's resources.  

Often we use type assertions to verify that you received a net.Error.
Type assertions provides access to an interface value's underlying concrete value.

```go
{
	var i interface{} = "hello"

	s := i.(string) //output - hello
	fmt.Println(s)

	s, ok := i.(string) //output - hello true
	fmt.Println(s, ok)

	f, ok := i.(float64) //output - 0 false
	fmt.Println(f, ok)

	f = i.(float64) // panic: interface conversion:
	fmt.Println(f)  // interface {} is string, not float64
}
```
Asserting a net.Error to check whether the error was tmp.  
```go
if nErr, ok := err.(net.Error); ok && !nErr.Temporary() {return err}
```
###Why we use Timeout  
We want to keep the application predictable and user-friendly. If a user/client  
performs a Dial and needs to wait for a response before going any further. Then we   
are at the mercy of the OP to Timeout for us, and this takes time.  
If the Serves isn't responding we want to time out quickly and move on. One solution  
to explicitly define the time period we want to wait through the DialTimeout function.
