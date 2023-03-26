# Arrays and Other Datatypes

## Arrays
An array in Go is size immutable like in Java. Arrays use a similar syntax to Java. The syntax being:
```go
var a [5]int
```
Which initializes an array of the size of 5 ints. To access an element in an array is:
```go
a[4] = 9
```
To find the length of an array use the `len()` function

## Slices
Slices are more mutable than arrays and generally superior (similar to Arrays vs ArrayLists in Java). To make a slice, you would use the `make()` function:
```go
s := make([]string, 3)
```
Setting and getting objects from a slice are the same as an Array, but slices can have values appended to them and be easily copied.

## Structs
Structs are very similar to Structs in C, where they are a collection of typed fields. The way to create a struct is:
```go
type person struct {
	name string
	age  int
}

p := person{'Bob', 25}

age = p.age
name = p.name

np := person{name: "Robert", 26}
```
As seen above, you can access the fields within a struct in a similar way to public fields in Java and C.

