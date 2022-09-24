package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Printf("Size %d and capacity %d \n", len(s), cap(s))
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	//index out of range
	// s[3] = "test"
	s = append(s, "test")
	fmt.Printf("Size %d and capacity %d \n", len(s), cap(s))
	fmt.Printf("Slice %v \n", s)
	s = append(s, "e", "f")
	fmt.Printf("Slice %v \n", s)

	c := make([]string, len(s))
	fmt.Printf("Before copy %s \n", c)
	copy(c, s)
	fmt.Printf("After copy %s \n", c)

	//return slices
	l := s[2:5]
	fmt.Println("sl1:", l)

	//return first 3
	l = s[:3]
	fmt.Println("sl2:", l)

	//skip first 2
	l = s[2:]
	fmt.Println("sl3:", l)

	//single dimension slice
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		twoD[i] = make([]int, i+1)
		for j := 0; j < len(twoD[i]); j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	//initialize 2d slice
	numbers := [][]int{{1, 2, 3}, {1}}

	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers[i]); j++ {
			fmt.Printf("%d ", numbers[i][j])
		}
		fmt.Println()
	}
}
