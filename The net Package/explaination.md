# The net Package
*From [pkg.go.dev/net](https://pkg.go.dev/net)*  
Package net provides a portable interface for network I/O, including TCP/IP, UDP, domain name resolution, and Unix domain sockets.
## The HTTP sub-package
The HTTP sub-package (net/http) allows for both HTTP client and server implementations.   
For this class, we will mostly deal with HTTP GET and POST requests.
```go
response, error := http.Get('http://www.google.com/')

response, error := http.Post('http://www.wpi.edu/')
```
To create a request with headers, you need some more configuration.
```go
client := &http.Client{}

resp, err := client.Get("http://example.com") // Without headers

req, err := http.NewRequest("GET", "http://example.com", nil) // Create the request to add headers an other information
req.Header.Add("If-None-Match", `W/"wyzzy"`) // Adding a header and value
resp, err := client.Do(req) // Actually send the request off
```

Finally, to read the body from a response, you would use: `body, err := io.ReadAll(resp.body)`

## What is RPC and the RPC sub-package

RPC or Remote Procedure Call is a form of Inter-Process communication that passes messages between a server and client in order to create a channel between the two. Essentially, the server exposes a method over the network for clients to use and call.  
net/rpc allows for the creation of both RPC clients and servers using Go, which is used in the MapReduce project. The RPC is widely used since it is able to detach the communcation and implementation. This means that each party doesn't care about the underlying code that is processed, only the inputs and outputs.