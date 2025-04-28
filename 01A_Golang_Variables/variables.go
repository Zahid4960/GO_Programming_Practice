package main

import "fmt"

// package level variable
var x int = 10

// at package level variable can be declared and initialized in one line
// var y := 20

var (
	id     float32 = 20047
	name   string  = "Zahid Hasan"
	gender string  = "male"
)

func main() {
	var a int // declaration
	a = 20047 // initialization

	var b int = 20047 // declaration and initialization both in one line

	c := 20047 // short declaration and initialization

	d := 3.50

	var e float32 = 3.50

	fmt.Println("Value of a is:", a)
	fmt.Println("Value of b is:", b)
	fmt.Println("Value of c is:", c)
	fmt.Printf("%T %v \n", d, d) // float64 3.5
	fmt.Printf("%T %v \n", e, e) // float32 3.5
	fmt.Println("Value of x is:", x)
	fmt.Println("My id is:", id)
	fmt.Println("My name is:", name)
	fmt.Println("My gender is:", gender)
}
