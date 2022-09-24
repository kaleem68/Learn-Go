package main

import (
	"fmt"
	"time"
)

var println = fmt.Println

func printTo15() {
	for i := 1; i <= 15; i++ {
		println("Fun1 ", i)
	}
}
func printTo10() {
	for i := 1; i <= 10; i++ {
		println("Fun2 ", i)
	}
}

// concurrency / go routine
func main() {
	go printTo15()
	go printTo10()
	time.Sleep(3 * time.Second)
}
