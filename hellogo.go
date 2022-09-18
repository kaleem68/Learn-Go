package main

import (
	"fmt"
)

var println = fmt.Println

func main() {
	slices()
}
func slices() {
	//var name []datatype
	sl1 := make([]string, 3)
	sl1[0] = "Mon"
	sl1[1] = "Tue"

	println("Length of slice ", len(sl1))

	for _, v := range sl1 {
		println(v)
	}
	// A slice points at an array and you can create a slice
	// of an array (A slice is a view of an underlying array)
	// You can have multiple slices point to the same array
	aArr := [5]int{10, 20, 30, 40, 50}
	// Start at 0 index up to but not including the 2nd index
	sl3 := aArr[0:2]
	println("Slices with first two elements ", sl3)

	//first two items
	println(aArr[:2])

	//last two items
	println(aArr[len(aArr)-2:])

	//changing the slice changes the array
	sl3[0] = -1
	println(aArr)

	// similarly changing the array changes the slice
	aArr[0] = 21
	println("slice ", sl3)
	println("array ", aArr)

	//appending value to slice overwrites the value in the array
	sl3 = append(sl3, 100)
	println("append modified array ", aArr)
	println("appended slice ", sl3)
	sl3 = append(sl3, 101)
	sl3 = append(sl3, 102)
	sl3 = append(sl3, 103)
	sl3 = append(sl3, 104)
	sl3 = append(sl3, 105)
	sl3 = append(sl3, 106)
	sl3 = append(sl3, 107)

	// append slice beyond array limit
	println("slice ", sl3, len(sl3))
	// appending the slice does not grow the array limits
	println("array ", aArr, len(aArr))

	// Printing empty slices will return nils which show
	// as empty slices
	emptySlices := make([]string, 6)
	println("empty slice ", emptySlices[0])
	println("emptySlices[0] :", emptySlices[0])

}
