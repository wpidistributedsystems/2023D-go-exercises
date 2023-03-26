# [Concurrency, Channels, and WaitGroups. Oh my!](https://youtu.be/J591SioKZJc?t=18)

## The `go` keyword
The go keyword is the way to spin off a concurrent process or goroutine. The keyword is used as follows:
```go
go func(arguments){
    // Do something then return
}("parameter")
```
In the example above, an anonymous function was used to task the goroutine. You can pass arguments by specifying the variable name to be used in the parenthesis next to the `func` and the argument to pass in the parenthesis next to the `}`

The `go` keyword can also be used with a concrete function as well.
```go
go fmt.Println("Hello!")
```

But wait, how do you communicate between goroutines?

## Channels
A channel is the datatype used to communicate between goroutines. There are two types of channels, unbuffered and buffered. A channel is made as follows:
```go
channel := make(chan int)
```

To read and write from channels, you need to use a specific operator `<-`
Unlike many other operators (`+`, `-`, and `*`), it is not communicative.

To write to a channel, we would place the operator like this:
```go
channel <- "variable you want to write"
```

To read from a channel, you would use it like this:
```go
a := <- channel
```
This operator blocks if the channel is empty and the routine would wait till it has something to read from the channel

**Notice that the location of the `<-` operator is very important** 

The main difference between a channel and a buffered channel is that a buffered channel can only have a set amount of objects within it before writing to it blocks.
A buffered channel is made like this:
```go
bufferedChannel := make(chan int, 8)
```
Where the size of the buffer is 8

## WaitGroups

A WaitGroup essentially acts as a way for all the goroutines to tell the main goroutine that they have finished processing. An example of these being used is in [a more complex example.go](./a%20more%20complex%20example.go) on lines 41 to 68  
To define a WaitGroup, you first need to import the "sync" package. Secondly, you would start with the definition: `var group sync.WaitGroup`. The waitgroup does not need to be initialized. Thirdly, in the goroutine, you need to call `group.Add(1)`. This tells the WaitGroup that there is a thread running it needs to block for. At the end of the goroutine, you need to add a call to `group.Done()` which tells the group there is one less goroutine to wait on.
Finally, in the controlling goroutine, you need to call `group.Wait()` which blocks the caller until the count of working threads is 0.
