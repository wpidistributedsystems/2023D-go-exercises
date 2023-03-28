# Go warmup

Go is a statically typed compiled programming language, similar to C.

It has some differences to C such as memory safety, a garbage collector, and structural typing. Additionally, Go is meant for concurrency, with its motto for concurrency being:  
`Do not communicate by sharing memory; instead, share memory by communicating`

Additionally, Go is white-space significant and does not need terminating characters on each line.

The recommended order is:
* [Primitive Types and Variables](./Variables/explaination.md)
* [Conditionals and Loops](./Conditionals/explaination.md)
* [Arrays and Other Datatypes](./Arrays%20and%20Other%20Datatypes/explaination.md)
* [Concurrency, Channels, and Waitgroups](./Concurrency/explaination.md)
* [The net Module](./The%20net%20Package/explaination.md)
There are exercises with some of the modules. It is recommended to do them, as the test file runs sequentially

## Running the Tests
We recommend using the docker-compose file in this repo. To run it install the latest version of docker and run `docker compose up --build` to rebuild and run the tests.

Alternatively, if you have installed the latest go version, you can also run `go test -v` to rebuild and run the tests. 