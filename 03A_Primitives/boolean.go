package main

import "fmt"

func main() {
	var n bool = true
	var m bool = false

	var x bool // by default, it is false

	fmt.Printf("%v %T \n", n, n) // true bool
	fmt.Printf("%v %T \n", m, m) // false bool
	fmt.Printf("%v %T \n", x, x) // false bool
}
