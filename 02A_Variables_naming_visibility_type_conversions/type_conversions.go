package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int = 10
	fmt.Printf("%T %v \n", a, a)

	var b float32
	b = float32(a) // Type conversion from int to float32`
	fmt.Printf("%T %v \n", b, b)

	var c string
	c = strconv.Itoa(a) // Type conversion from int to string
	fmt.Printf("%T %v", c, c)
}
