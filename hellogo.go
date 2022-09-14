package main

import (
	"fmt"
	"time"
)

var out = fmt.Println

func main() {
	timeTest()
}
func timeTest() {
	now := time.Now()
	// Get day, month, year and time data
	out(now.Day(), now.Month(), now.Year())
	out(now.Hour(), now.Minute(), now.Second())
}
