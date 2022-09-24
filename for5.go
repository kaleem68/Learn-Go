package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		fmt.Printf("%d ", i)
		i++
	}
	fmt.Println()

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	i = 1
	for {
		if i > 10 {
			break
		}
		fmt.Printf("%d ", i)
		i++
	}
	fmt.Println()

	//skip even
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
	println()

}
