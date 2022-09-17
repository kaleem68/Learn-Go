package main

import (
	"fmt"
)

var println = fmt.Println

func main() {
	arrays()
}
func arrays() {

	for i := 1; i <= 10; i++ {
		println(i)
	}
	println()

	//iterate 10 to 1
	for i := 10; i >= 1; i-- {
		println(i)
	}
	println()

	//using for loop as a while loop

	count := 1
	for count <= 10 {
		println(count)
		count++
	}
	println()

	// ----- ARRAYS -----
	// Collection of values with the same data type
	// and the size can't be changed
	// Default values are 0, 0.0, false or ""

	//Arrays
	aNums := []int{1, 2, 3}
	for index, num := range aNums {
		println(index, num)
	}
	println()
	//skip index with _
	for _, num := range aNums {
		println(num)
	}
	println()

	//create and display array
	var arr1 [5]int
	println(arr1)

	println()
	arr1[0] = 1
	println(arr1)

	//declear and initialize
	arr2 := [5]int{1, 2, 3, 4, 5}

	//simple iteration
	for i := 0; i < len(arr2); i++ {
		println(arr2[i])
	}
	println()
	//iterate with range
	for i, v := range arr2 {
		fmt.Printf("index: %d, value: %d ", i, v)
	}
	println()
	//multidimensional array
	arr3 := [2][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}
	for row := 0; row < 2; row++ {
		for col := 0; col < 5; col++ {
			fmt.Printf("%d: ", arr3[row][col])
		}
		println()
	}

	println("Iterate 2d array")
	//iterate 2d array with row length and col length
	for row := 0; row < len(arr3); row++ {
		for col := 0; col < len(arr3[row]); col++ {
			fmt.Print(arr3[row][col], " ")
		}
		println()
	}
	println()
	//multidimensional array without column mentioned
	arr4 := [2][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8},
	}
	println(arr4)

}
