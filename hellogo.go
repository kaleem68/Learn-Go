package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var out = fmt.Println

func main() {
	mathTest()
}
func mathTest() {
	out("Simple maths")
	out("5 + 4 ", (5 + 4))
	out("5 - 4 ", (5 - 4))
	out("5 * 4 ", (5 * 4))
	out("5 / 4 ", (5 / 4))
	out("5 % 4 ", (5 % 4))

	out("Increment")
	count := 1
	count++
	count += 1
	out(count)

	//Float precision
	out("Precise ", 0.11111111111111111111111111111111111111111111+0.11111111111111111111111111111111111111111111)

	//random number between 1 - 50
	//since 01/01/1970
	timeSeconds := time.Now().Unix()
	rand.Seed(timeSeconds)
	randomNumber := rand.Intn(50) + 1
	out("Random digit ", randomNumber)

	//math functions
	out("Abs(-10) =", math.Abs(-10))
	out("Pow(4, 2) =", math.Pow(4, 2))
	out("Sqrt(16) =", math.Sqrt(16))
	out("Cbrt(8) =", math.Cbrt(8))
	out("Ceil(4.4) =", math.Ceil(4.4))
	out("Floor(4.4) =", math.Floor(4.4))
	out("Round(4.4) =", math.Round(4.4))
	out("Log2(8) =", math.Log2(8))
	out("Log10(100) =", math.Log10(100))
	// Get the log of e to the power of 2
	out("Log(7.389) =", math.Log(math.Exp(2)))
	out("Max(5,4) =", math.Max(5, 4))
	out("Min(5,4) =", math.Min(5, 4))

	// Convert 90 degrees to radians
	r90 := 90 * math.Pi / 180
	// Convert 1.5708 radians to degrees
	d90 := r90 * (180 / math.Pi)

	fmt.Printf("%f radians = %f degrees \n", r90, d90)

	//trig
	out(math.Sin(90))

	// There are also functions for Cos, Tan, Acos, Asin
	// Atan, Asinh, Acosh, Atanh, Atan2, Cosh, Sinh, Sincos
	// Htpot

	//Formatted Print, inspired by C

	/*
		%d -> Integers
		%c -> Character
		%f -> Floats
		$t -> Boolean
		%s -> String

		%o -> Base 8
		%x -> Base 16
		%v -> Guess based on input

		%T : Gives Type of supplied value

	*/
	fmt.Printf("%d, %c, %f, %t, %s, %o, %x %v \n", 21, 'A', 3.14, true, "Ahan", 21, 172, "Guess!")
	fmt.Printf("Type: %T \n", true)

	//Float formatting
	fmt.Printf("no precision: %.f \n", 3.142)
	fmt.Printf("round two decimal: %.2f \n", 3.142)

	fmt.Printf("9 spaces: %9f \n", 3.14)
	fmt.Printf("9 spaces 2 precision: %9.2f \n", 3.14)

	// Sprintf returns a formatted string instead of printing
	sp1 := fmt.Sprintf("%9.f", 3.14)
	out(sp1)

}
