package main

import (
	date "example/project/mypackage"
	"fmt"
	"log"
)

var println = fmt.Println

func main() {
	date := date.MyDate{}
	dateError := date.SetDay(14)
	if dateError != nil {
		log.Fatal(dateError)
	}
	monthError := date.SetMonth(1)
	if monthError != nil {
		log.Fatal(monthError)
	}
	yearError := date.SetYear(2020)
	if yearError != nil {
		log.Fatal(yearError)
	}
	println(date.Day(), date.Month(), date.Year())
}
