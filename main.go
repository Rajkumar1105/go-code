package main

import "fmt"

// interfaces --have some kind of contract or some kind of agreement with the objects
type Humans struct {
	Name string
}

func (h Humans) Angry() {
	fmt.Println(h.Name, "is angry")
}

func (h Humans) Speak() {
	fmt.Println(h.Name, "is speaking in telugus")
}

func (h Humans) Hunt() {
	fmt.Println(h.Name, "hunting mosquitos")
}
func (h Humans) Play() {
	fmt.Println(h.Name, "is playing crickte")
}
func (h Humans) Eat() {
	fmt.Println(h.Name, "is easting hyderbadi birayni")
}

type Cat struct {
	Tail int
}

func (h Cat) Angry() {
	fmt.Println("is angry")
}

func (h Cat) Speak() {
	fmt.Println("is roaring")
}

func (h Cat) Hunt() {
	fmt.Println("cat would be eayting rabbits or deers")
}
func (h Cat) Play() {
	fmt.Println("Playing  with tigers")
}
func (h Cat) Eat() {
	fmt.Println("will eat dear")
}

// the whole purpose of an interface is to give the generallized behaviour for concrete objects
type Animal interface {
	Angry()
	Speak()
	Eat()
	Play()
	Hunt()
	Crawl(p int) string
}

// interface{}
func Party(p interface{}) {
	fmt.Println(p)
}

// interface{} is a super type and every object and every built in type satisfies it
func main() {

	mahesh := Humans{
		Name: "Mahesh",
	}

	mahesh.Eat()
	mahesh.Angry()
	mahesh.Hunt()
	mahesh.Speak()
	mahesh.Play()

	// We can say mahesh is animal

	//tiger := Cat{
	//	Tail: 1,
	//}

	Party(3)
	Party("ssay")
	Party(mahesh)

	mc := McDonalds{
		Location: "Hyderabad",
		Menu:     "Burger,MMc Maharaja ,alootokki",
	}

	//fmt.Println(mc.GetMenu())

	Operator(mc)
}

func Operator(fr Franchise) {
	fmt.Println(fr.PrepareFood())
	fmt.Println(fr.GetLocation())
	fmt.Println(fr.GetMenu())
}

type Data2 struct {
}

/**
   Functional Requirements
	1.	Define an interface named Franchise that declares the following methods:
	•	PrepareFood() string
	•	GetMenu() []string
	•	GetLocation() string
	2.	Create at least three different franchises (e.g., McDonald’s, KFC, Burger King) that implement the Franchise interface.
	•	Each franchise must return:
	•	a unique menu
	•	a unique food preparation result
	•	a different location
	3.	Implement a function named Operate(fr Franchise) that:
	•	Accepts a Franchise interface
	•	Prints the franchise location
	•	Prints the menu
	•	Prepares food

Type Franchise interface{
   PrepareFood() string
	•	GetMenu() []string
	•	GetLocation() string

}
*/

type Franchise interface {
	PrepareFood() string
	GetMenu() string
	GetLocation() string
}

type McDonalds struct {
	Menu     string
	Location string
}

func (m McDonalds) PrepareFood() string {
	return "Preparing food"
}

func (m McDonalds) GetMenu() string {
	return m.Menu
}

func (m McDonalds) GetLocation() string {
	return m.Location
}

type KFC struct {
	Menu    string
	Addrses string
}

func (m KFC) PrepareFood() string {
	return "Preparing food"
}

func (m KFC) GetMenu() string {
	return m.Menu
}

func (m KFC) GetLocation() string {
	return m.Addrses
}
