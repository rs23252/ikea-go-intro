package main

func main() {
	/* var i = 0
	for i = 1; i < 11; i++ {
		println(i)
	} */

	/* var i = 10
		for i > 0 {
	  println(i)
	  i--
	} */

	/* // Go uses the `ASCII` codes within the `for` loop to print alphabets from `a` to `z`.
	for ch := 'a'; ch <= 'z'; ch++ {
		fmt.Printf("%c ", ch)
	} */
	st := ""
	for i := 0; i < 5; i++ {
		st = st + "* "
		println(st)
	}
}
