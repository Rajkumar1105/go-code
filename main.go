package main

import (
	"fmt"
	"reflect"
)

type (
	// ID is user-defined data type
	ID uint32
	// Name is user-defined data type
	Name string
	// Surname is user-defined data type
	Surname string
)

// currency is my own custom type or we can call it as user defined type
type Currency int32

func convert(g Currency) {
	fmt.Println(g)
}

type Human struct {
	Legs      int
	Fingers   int
	HairColor string
}

func NewHuman(legs, finger int, haircolor string) Human {
	return Human{
		Legs:      legs,
		Fingers:   finger,
		HairColor: haircolor,
	}
}

// This is a method
func (h Human) Walk() { // h is  a receiver
	fmt.Println("hello world")
}

type Circle struct {
	Radius int
}

func main() {
	var g Currency
	fmt.Println(reflect.TypeOf(g))
	convert(g)
	//g = 90
	//fmt.Println(g)
	//
	//var name Name
	//name = Name("shubham")

	// type conversion
	g = 9
	fmt.Println(reflect.TypeOf(g))
	// we can do type conversion and we can convert it to float
	h := float32(g) // type conversion
	fmt.Println(reflect.TypeOf(h))

	//eswar := Human{
	//	Legs:      2,
	//	Fingers:   20,
	//	HairColor: "Black",
	//}

	hrithik := NewHuman(2, 21, "blonde")

	fmt.Println(hrithik.Fingers)

	//// anonymous
	//d := struct {
	//	Name string
	//}{
	//	Name: "shubham",
	//}

	//var c Circle
	//c.Radius = 9
	//fmt.Println(c)
	//
	hrithik.Walk()
}

// ----go is not an object oriented progrmaming language it has objects
// object is just a real world entity which we represent in our code
/// - how an object --what object contains ---> object has properties and object has behaviour
