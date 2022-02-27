## Slices
A slice is an ordered collection of values of a single type. The syntax for declaring a slice variable is very similar to declaring a scalar variable.

```go
var i int   // an int called i
var j []int // a slice of ints called j
```
In this example,
- i is a variable of type `int`.
- j is a variable of type `[]int`, that is, a `slice` of `int`.

Slices are very important in Go programs

Note: A slice is not an array. Go also supports arrays, but you'll see later than they aren't
very common, or very easy to use, so we won't discuss them at the moment.

## How large is a slice?

If I declare a slice, `[]string`, how many items can it hold?

The zero value of a slice is empty, that is, it has a length of zero; it can hold 0 items.

```go
package main

import "fmt"

func main() {
  var s []string
  fmt.Println(len(s))
}
```
We can retrieve the length of a slice with the built-in `len` function.

## Making a slice
We can create a slice with space to hold items using the built-in make function.

```go
package main

import "fmt"

func main() {
  var cities []string
  cities = make([]string, 20)
  fmt.Println(len(s))
}
```

In this example, on the first line `var cities []string` declares cities to be a slice of string. On the second line, cities is assigned the result of `make([]string, 20)`.

Lab:
Let's do a quick exercise to familarise yourself with using slices

```go
package main

import "fmt"

func main() {
  // Declare a variable called i which is a slice of 5 int.
  // Declare a variable called f which is a slice of 9 float64.
  // Declare a variable called s which is a slice of 4 string.

  fmt.Println(len(i), len(f), len(s))
}
```

- Does your program print the expected result, 5 9 4?
<details>
  <summary>Not sure how?</summary>

  ```go
    func main() {
      var i []int
      var f []float64
      var s []string
      
      i = make([]int, 5)
      f = make([]float64, 9)
      s = make([]string, 4)

      fmt.Println(len(i), len(f), len(s))
    }  
  ```
</details>
<br>