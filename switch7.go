package main

import (
	"fmt"
	"time"
)

func main7() {

	day := "Wed"
	basicSwitch(day)
	println()

	//multiple cases
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Printf("Its weekend %s \n", time.Now().Weekday())
	default:
		fmt.Printf("Its weekday %s \n", time.Now().Weekday())
	}

	//no switch value: alternative of if else
	switch {
	case time.Now().Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
	//use switch to find type

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	fmt.Println()
	whatAmI(false)
	whatAmI(23)
	whatAmI("truth")
	fmt.Println()
	whatAmIFunc(false)
	whatAmIFunc(23)
	whatAmIFunc("truth")

}
func whatAmIFunc(i interface{}) {
	switch t := i.(type) {
	case bool:
		fmt.Println("I'm a bool")
	case int:
		fmt.Println("I'm an int")
	default:
		fmt.Printf("Don't know type %T\n", t)
	}

}
func basicSwitch(day string) {
	switch day {
	case "Mon":
		fmt.Printf("Its Mon")
	case "Tue":
		fmt.Printf("Its Tue")
	case "Wed":
		fmt.Printf("Its Wed")
	case "Thur":
		fmt.Printf("Its Thur")
	case "Fri":
		fmt.Printf("Its Fri")
	case "Sat":
		fmt.Printf("Its Sat")
	default:
		fmt.Printf("Its Sund")
	}

}
