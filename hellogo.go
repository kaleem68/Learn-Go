package main

import (
	"fmt"
	"unicode/utf8"
)

var out = fmt.Println

func main() {
	runes()
}
func runes() {
	// In go characters are called runes
	// Runes are unicodes that represent characters
	tStr := "abcdefg"
	out("Rune Count : ", utf8.RuneCountInString(tStr))

	for i, runeVal := range tStr {
		//Getindex, Rune unicode, character
		fmt.Printf("%d : %#U, %c\n", i, runeVal, runeVal)
	}
}
