package main

import (
	"fmt"
	"math"
	"reflect"
)

const s1 = "adsad"
const s2 string = "test2"

func main4() {
	fmt.Printf("s1: %v type %v \n", s1, reflect.TypeOf(s1))
	fmt.Printf("s2: %v type %v \n", s2, reflect.TypeOf(s2))

	const n = 100
	const d = 1e20 / n
	fmt.Println(d, reflect.TypeOf(d))

	fmt.Println(math.Sin(n))

}
