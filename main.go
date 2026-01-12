package main

import "fmt"

type Data struct {
	Name string
}

//func (d *Data) GetName() string {
//	return d.Name
//}

func (v *Data) SetName(name string) {
	v.Name = name
}

// lets say db connectioio   App ---->Db connection

func main() {
	//c := Data{
	//	Name: "shubham",
	//}

	// new keyword creates the object and returns pointer

	c := new(Data)

	c.SetName("Shubham")

	fmt.Println(c)

	// new and make
	//c.SetName("John")
	//fmt.Println(c)

	//var f *Data
	//
	//f = &c
	//
	//fmt.Println(f) // f is a pointer to struct
	//
	//fmt.Println((*f).Name) // so first dereference the struct  then we are accessing the values

}
