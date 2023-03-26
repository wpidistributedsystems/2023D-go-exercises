# Primitive types and variables

## Types
There are 4 different primitive types:  

| Type     | Examples and Syntax |
|----------|---------------------|
| Integers | 1                   |
| Floats   | 3e20, 3.14          | 
| Booleans | true, false         |
| Strings  | "Hello!"            |

These values can be combined in similar ways to Java:

* int + float = float
* int + int = int
* string + string = string
* bool || bool = bool

## Variables
The `var` keyword is used to declare variables. Unlike C and Java, the type is inferred from the initialized value.  
The line: `var a = 10` implicitly sets the type of `a` to 10.  
Variables can also be declared and initialized by using the walrus (`:=`) operator. The type of variables set this way will be inferred as well.  
Multiple variables can be set on one line separating the variable names and values by commas, like this:  
`a, b := 5, 10`  
The same syntax can be used with the `var` keyword.  
Finally, variables can be initialized to a "zero-value" by using this syntax: `var a int`. The variable a would be initialized to a value of 0.  
