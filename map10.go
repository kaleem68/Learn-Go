package main

import (
	"fmt"
)

func main10() {
	dldMarks := make(map[string]int)
	dldMarks["aman"] = 18
	dldMarks["kaleem"] = 29

	for key, val := range dldMarks {
		fmt.Printf("key %s, val %d \n", key, val)
	}
	delete(dldMarks, "kaleem")
	_, present := dldMarks["kaleem"]
	fmt.Println("kaleem marks present", present)

	fullNames := map[string]string{"kaleem": "kaleemullah", "andrew": "josh andrew"}
	fmt.Println(fullNames)

	defaultMap := make(map[string]int)
	defaultMap["default"] = 21
	fmt.Println(defaultMap)

	//error: assignment to entry in nil map
	// var defaultMap2 map[string]int
	// defaultMap2["default"] = 21
	// fmt.Println(defaultMap2)

}
