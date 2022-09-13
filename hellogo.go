package main

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
)

var pl = fmt.Println

func main() {
	casting()
}
func casting() {
	cV1 := 1.7
	cV2 := int(cV1)
	pl(cV2)

	cV3 := "342"
	// Atoi = ASCII to integer
	cV4, err := strconv.Atoi(cV3)
	if err == nil {
		pl(cV4, reflect.TypeOf(cV4))
	} else {
		log.Fatal(err)
	}

	//Itoa = Integer to ASCII
	cv5 := strconv.Itoa(32)
	pl(cv5, reflect.TypeOf(cv5))

	//string to float
	cV6 := "3.1415"
	floatCV6, err2 := strconv.ParseFloat(cV6, 64)

	if err2 == nil {
		pl(floatCV6, reflect.TypeOf(floatCV6))
	} else {
		log.Fatal(floatCV6)
	}

	//string to float cont
	pie := "3.14"
	if res, err3 := strconv.ParseFloat(pie, 64); err3 == nil {
		pl(res, reflect.TypeOf(res))
	} else {
		log.Fatal(err3)
	}

	//float to string: formatting similar to c
	pieStr := fmt.Sprintf("%f", 3.14)
	pl(pieStr, reflect.TypeOf(pieStr))

}
