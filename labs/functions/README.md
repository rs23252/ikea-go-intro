## Functions
Go is all made up of functions, in fact, you've been writing functions all along.

You can declare your own functions with the func declaration.

A function's name must be a valid identifier, just like `const` and `var`.

```go
func print(msg string) {
  fmt.Println(msg)
}

func main() {
  print("Hello World!!")
}
```

## Parameters
Parameters are defined in the function definition and the argument passed to the function should match the data type.

```go
// here the msg is the variable name and string is the datatype this function expects as a parameter
func print(msg string) {
  fmt.Println(msg)
}

hello(1234)
// compile this code and see that happens
```

## Return value
The function definition should define the return type if the method returns a value.

```go
// Here last int is the return type
func sum(a int, b int) int {
  return a + b
}
```

One of Go's unusual features is that functions and methods can return multiple values. In case of multiple return type, you need to put them in parenthesis (n int, err error)

```go
// In Go, Write function can return a count and an error: "Yes, you wrote some bytes but not all of them because the disk was full".

func Write(b []byte) (n int, err error)
```

## Named result parameters
The return or result "parameters" of a Go function can be given names. 

The names are not mandatory but they can make code shorter and clearer: they're documentation. If we name the results of `nextInt` it becomes obvious which returned `int` is which.

```go
func nextInt(b []byte, pos int) (value, nextPos int)
```

## Defer
Go's defer statement schedules a function call (the deferred function) to be run immediately before the function executing the defer returns. It's an unusual but effective way to deal with situations such as resources that must be released regardless of which path a function takes to return. The canonical examples are unlocking a mutex or closing a file. 

```go
func main() {
  defer releaseResources()
  fmt.Println("testing defer statement")
}

func releaseResources() {
  fmt.Println("closing and releasinng connections....")
}
```
