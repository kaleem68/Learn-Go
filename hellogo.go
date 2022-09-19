package main

import (
	"fmt"
)

var println = fmt.Println

func main() {
	pointers()
}

func pointers() {
	// ----- POINTERS -----

	// You can pass by reference with the &
	// (Address of Operator)
	// Print amount and address for amount in memory
	f4 := 10
	println("f4 value ", f4)
	println("f4 address ", &f4)

	//store address of variable
	var f4Ptr *int = &f4
	println("f4Ptr address", f4Ptr)
	println("f4Ptr address value", *f4Ptr)
	*f4Ptr++
	println("f4Ptr new address value", *f4Ptr)

	pArray := [4]int{1, 2, 3, 4}
	doubleValuesPtr(&pArray)
	println(pArray)

}
func doubleValuesPtr(arrPtr *[4]int) {
	for i := 0; i < 4; i++ {
		arrPtr[i] *= 2
	}
}
