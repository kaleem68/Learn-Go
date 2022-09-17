package main

import (
	"fmt"
)

var println = fmt.Println

func main() {
	runes()
}
func runes() {
	// String into slice of runes
	aString1 := "abcde"
	rArr := []rune(aString1)

	for _, v := range rArr {
		fmt.Printf("%c ", v)
	}
	fmt.Printf("\n")

	// Byte array to string
	byteArr := []byte{'a', 'b', 'c'}
	bStr := string(byteArr[:])
	println("byte string ", bStr)
}
