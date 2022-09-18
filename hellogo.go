package main

import (
	"fmt"
	"log"
)

var println = fmt.Println

func main() {
	// ----- FUNCTIONS -----
	// func funcName(parameters) returnType {BODY}
	// If you only need a function in the current package
	// start with lowercase letter
	// Letters and numbers in camelcase
	functions()
}
func functions() {
	//function returning two values
	a, b := 10, 20
	println("a =", a, "b =", b)
	a, b = swap(a, b)
	println("a =", a, "b =", b)

	//function returning error
	ans, err := getQuotient(5, 1)
	if err == nil {
		println("Result ", ans)
	} else {
		// println(err)
		//End of program
		log.Fatal(err)
	}
	// Function receives unknown number of parameters
	// Variadic Function
	println("Sum of numbers: 31, 9, 10, 20, 20, 10 =", sumAll(31, 9, 10, 20, 20, 10))

	// Pass an array to a function by value
	oddNumbers := []int{1, 3, 5, 7, 9}
	println("Sum of odd: [1, 3, 5, 7, 9] =", sumArray(oddNumbers))

	//pass by ref: modify array value inside function
	println("before modify array[0]", oddNumbers)
	modifyArray(oddNumbers)
	println("after modify array[0]", oddNumbers)

	//pass by val: changing primitive inside function
	value := 5
	println("before modify value 5:", value)
	changeVal(value)
	println("after modify value 5:", value)

}
func swap(a, b int) (int, int) {
	return b, a
}
func getQuotient(x float64, y float64) (float64, error) {
	if y == 0 {
		//return error with a dummy value
		return 0, fmt.Errorf("Cannot be divide by zero")
	}
	return x / y, nil
}

func sumAll(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func sumArray(nums []int) (sum int) {
	sum = 0
	for _, num := range nums {
		sum += num
	}
	return
}
func modifyArray(nums []int) {
	nums[0] = 99
}

func changeVal(val int) int {
	val += 1
	return val
}
