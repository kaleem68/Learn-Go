package main

import (
	"fmt"
)

var println = fmt.Println

type Contact struct {
	fName, lName, phone string
}
type Business struct {
	name    string
	address string
	contact Contact
}

func (b *Business) DisplayInfo() {
	fmt.Printf("business name: %s, person name: %s, %s \n", b.name, b.contact.fName, b.contact.lName)
}

func main() {
	Composition()
}
func Composition() {
	contactOne := Contact{"Kaleem", "Niz", "0300"}
	BusinessOne := Business{"Software", "Remote", contactOne}
	BusinessOne.DisplayInfo()
}
