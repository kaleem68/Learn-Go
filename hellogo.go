package main

import (
	"fmt"
	"reflect"
)

var println = fmt.Println

type Animal interface {
	Eat()
	Sleep()
}

// cat
type Cat string

func (cat *Cat) Name() string {
	return string(*cat)
}

func (cat Cat) Eat() {
	println("Cat is eating")
}

func (cat Cat) Sleep() {
	println("Cat is Sleep")
}

// dog
type Dog string

func (dog *Dog) Bark() {
	fmt.Printf("%s is barking \n", dog.Name())
}
func (dog *Dog) Name() string {
	return string(*dog)
}
func (dog Dog) Eat() {
	println("Dog is eating")
}

func (dog Dog) Sleep() {
	println("Dog is Sleep")
}

func main() {
	var cat Animal = Cat("kitty")
	cat.Eat()
	cat.Sleep()
	fmt.Printf("type %s \n", reflect.TypeOf(cat))
	var newCat Cat = cat.(Cat)
	fmt.Printf("type %s, name, %s \n", reflect.TypeOf(newCat), newCat.Name())

	println()
	var dog Animal = Dog("tom")
	dog.Eat()
	dog.Sleep()
	fmt.Printf("type %s \n", reflect.TypeOf(dog))
	var newDog Dog = dog.(Dog)
	newDog.Bark()
	fmt.Printf("type %s, name, %s \n", reflect.TypeOf(newDog), newDog.Name())

}
