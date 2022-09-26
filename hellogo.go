package main

import (
	"fmt"
)

func IsPositive(num int) (bool, error) {
	if num > 0 {
		return true, nil
	}
	return false, fmt.Errorf("not a positive number")
}

func main() {
	x := 0
	fmt.Println("lol", float64(2/x))
}
