package main

import (
	"fmt"
	"reflect"
)

// Variadic functions: k argument
func variadic(values ...int) int {
	values[0] = 21
	fmt.Println("values ", reflect.TypeOf(values))
	return len(values)
}

// return multiple values
func vals() (int, int) {
	return 3, 7
}
func main() {
	//return multiple values
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)

	fmt.Println("length ", variadic(1, 2, 3, 4, 5))

	mySlice := []int{1, 2, 3, 4}
	fmt.Println()
	fmt.Println("slice before update", mySlice)
	fmt.Println("length ", variadic(mySlice...))
	fmt.Println("slice after update ", mySlice)

}
