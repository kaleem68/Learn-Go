package main

import "fmt"

func main6() {

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}
	if num := 19; num < 0 {
		fmt.Println("Negative ")
	} else if num < 10 {
		fmt.Println("Has 1 digit")
	} else {
		fmt.Println("Has multiple digits")
	}
}
