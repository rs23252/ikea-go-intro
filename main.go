package main

import "fmt"

func main() {
	t := sum(4, 5)
	fmt.Println(t)
}

func sum(a int, b int) int {
	total := a + b
	return total
}
