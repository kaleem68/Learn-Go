package main

import (
	"fmt"
)

var println = fmt.Println

func main() {
	mapBasics()
}

func mapBasics() {
	intMap := make(map[int]int)
	intMap[2] = 22
	fmt.Println(intMap[2])

}
