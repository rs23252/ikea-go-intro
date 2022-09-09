package main

import (
	"fmt"
	"sort"
)

func main() {
	/* // Declare a variable called i which is a slice of 5 int.
	var i [5]int
	// Declare a variable called f which is a slice of 9 float64.
	var f [9]float64
	// Declare a variable called s which is a slice of 4 string.
	var s [4]string

	fmt.Println(len(i), len(f), len(s)) */

	// first make a slice of strings with all the letters from a to z
	LetterValue := make([]string, 0)
	LetterValue = append(LetterValue, "")
	for pos := 'a'; pos <= 'z'; pos++ {
		char := fmt.Sprintf("%c", pos)
		LetterValue = append(LetterValue, char)
	}
	fmt.Println(LetterValue)
	sumofletters("bc", LetterValue)

}

func sumofletters(s string, LetterValue []string) {
	total := 0
	var lv string
	for i := 0; i < len(s); i++ {

		lv = s[i]
		total = total + sort.SearchStrings(LetterValue, lv)
	}
	fmt.Println(total)

}
