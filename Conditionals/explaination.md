# Conditionals and Loops

## If, Else, and Switch statements
If statements are essentially the same as Java, without the parentheses. They can optionally have an else or if else statement.  
However, there is **NO** ternary if.  
```go
if i {
	// evaluates if i is true
} else {
	// evaluates if i is false
}
```
Switch statements are denoted by the `switch` keyword and the cases by the `case` and `default` keywords, as shown in the block below:
```go
switch i {
    case 0:
            fmt.Println("Got a 0!")
    case 1:
	    fmt.Println("Got a 1!")
    case 2:
	    fmt.Println("Got a 2!")
    case 3, 4, 5:
		fmt.Println("Got a 3, 4, or 5!")
    default:
		fmt.Println("Not a 0, 1, or 2!")
}
```
You can also use switch statements to compare types of variables, using the syntax below:
```go
switch i.(type) {
    case bool:
		fmt.Println("i is a bool!")
    case int:
		fmt.Println("i is an int!")
    default:
		fmt.Printf("i is a %T\n", i.(type))
}
```

## For loops
```go
i := 1
for i <= 10 {
	fmt.Println(i)
	i = i + 1
}

for j := 2; j <= 10; j++ {
	fmt.Println(j)
}

for {
	fmt.Println("loop")
}
```  
For loops in Go can be done with a single condition, an inline condition like in Java, or without a condition.  
All of these different types of loops are shown in the block above.   
Loops also can use the `break` and `continue` keywords to break out of the loop or continue onto the next loop iteration.  

