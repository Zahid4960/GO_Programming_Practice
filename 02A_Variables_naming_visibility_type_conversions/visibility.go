package main

import "fmt"

var A int = 10 // Global level variable
var a int = 20 // Package level variable

func main() {
	var c float32 = 3.50             // Local variable
	fmt.Println("Value of A is:", A) // 10
	fmt.Println("Value of a is:", a) // 20
	fmt.Println("Value of c is:", c) // 3.5
}
