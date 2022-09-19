package main

import (
	"fmt"
)

var println = fmt.Println

func main() {
	pointers()
}

func pointers() {
	// Pointers: basics
	myValue := 10
	println("value", myValue)
	println("address", &myValue)
	var valuePtr *int = &myValue
	println("address ", valuePtr)

	println("before change ", myValue)
	changeValuePtr(&myValue)
	println("after change ", myValue)

}

func changeValuePtr(valuePtr *int) {
	*valuePtr += 1
}
