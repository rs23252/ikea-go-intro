## Unit Testing

Testing is built in to the language and the syntax is same as you write rest of the code.

Writing a test is just like writing a function, with a few rules:
- Test needs to be in a file with a name like `xxx_test.go`
- The test function must start with the word `Test`
- The test function takes one argument only `t *testing.T`
- In order to use the `*testing.T` type, you need to import `"testing"` package

For now, it's enough to know that `t` of type `*testing.T` is your "hook" into the testing framework so you can do things like `t.Fail()` or `t.Error()` when you want to fail.

A simple test in Go:

```go
package main

import (
  "testing"
)

func TestGreeting(t *testing.T) {
  got := greeting()
  want := "Hello"
  if got != want {
    t.Errorf("greeting = %s; wanted Hello", got)
  }
}

func greeting() string {
  return "Hello"
}
```

## Test coverage
`go test` can report coverage

```shell
go test -coverprofile=cover.out
```

This produces a coverage file, `cover.out`:

- `go tool cover -func=cover.out` will print the coverage report
- `go tool cover -html=cover.out` will open the report in a browser