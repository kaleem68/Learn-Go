package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	strings()
}
func strings() {
	sV1 := "A word"
	pl(len(sV1))
}
