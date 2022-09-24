package main

import (
	"fmt"
)

var println = fmt.Println

type Person struct {
	name string
	age  int
}

func (p *Person) SetName(pName string) {
	p.name = pName
}
func (p *Person) SetAge(pAge int) {
	p.age = pAge
}

func (p *Person) GetName() string {
	return p.name
}
func (p *Person) GetAge() int {
	return p.age
}

func main() {
	BasicStruct()
}
func BasicStruct() {
	kaleem := Person{"kaleem", 22}
	fmt.Printf("name %s, age %d \n", kaleem.GetName(), kaleem.GetAge())
	kaleem.SetName("Kaleemullah")
	kaleem.SetAge(23)
	fmt.Printf("name %s, age %d \n", kaleem.GetName(), kaleem.GetAge())
}
