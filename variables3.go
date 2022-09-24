package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	variable()
}
func variable() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

	var p Person
	fmt.Println("Default struct ", p)

}
