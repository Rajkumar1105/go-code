package main

import "fmt"

// g:=9  // cannot do this outside the function
var g = 9 // global variable

// define multiple variable definitions
var (
	a int    = 3
	x string = "Hello world"
)

const alpha = 90 // once they are declared you cannot change them

func main() {

	//var c bool
	//var c string // uniocde (utf-8),asciii

	/**
	int  int8  int16  int32  int64  // it can store both positivce and negative
	uint uint8 uint16 uint32 uint64 uintptr  // it can only store insiogned values that is positive values
	*/

	//var z int32

	//var z uint8 // it can store an integer that can be respreented in 8 bits
	//z = -127    // the maximum value that it cna store is from -127 to 127
	//fmt.Println(z)
	// uintptr is for storing pointer variables
	//var d uint8
	//d = 9
	//var c uintptr
	//c = d

	// byte // 1 byte =8 bits
	//
	//var c byte  // []byte slice
	//
	//c = 9
	// rune  // alias for int32
	//var c rune
	//c = 90

	//float32 float64

	//var t float32
	//t = 8999999999999999999999999999999999999999.900000000000000000000000000

	// complex64 complex128 // 2+3i (-1)
	//var z1 complex128 = 3 + 4i
	//var z int8 = 2 // [10] ---> // [00000010]
	//fmt.Println(z)

	//var s float32
	////fmt.Println(s)
	////var z =9
	////var x =&z  // x is basically a pointer  // default valu of pointer is nil
	//
	////var t *int
	////fmt.Println(t)
	//s = 90
	//fmt.Println(s)

	//var c int
	//c = 9 // assignment opertaor
	// shorthand notation
	//c:=9 // compiler is smart enough to understand that the data type of c is 9
	//var g = 9 // the compiler is smart enough to understand the type

	var e = 5.5
	var d = false

	fmt.Println(d, e)
	a = 90

	f := Sum(9, 8) // arguments
	fmt.Println(f)
}

// You have to setup a go app and with   one folder and it  should be called as
// directions it will have a file called called direction.go and it will contain all the const values
// of directions for ex const North=1 and so on
// print these directions in main.go file

// void ff
// A function is a set of encapsulated code that you can re use as many times as you want

func Sum(a int, b int) (int, int, int) { // declaration and definition
	// is where logic goes
	sum := a + b
	//fmt.Println(sum)
	return sum, a - b, a * b
}

// parameter and arguments

// function call // where you are actually calling the function

// Now create a function that takes 5 int parameters   in the package called as fun paclage and the file would fun.go inside the package
// Now calculate the sum the difference and the product of all five numbers and return it
// you have to call it inside the main.go
