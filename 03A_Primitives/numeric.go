package main

import "fmt"

func main() {
	n := 30

	var a int8 = 10
	var b int16 = 20
	var c int32 = 40
	var d int64 = 50

	var w uint8 = 10
	var x uint16 = 20
	var y uint32 = 40
	var z uint64 = 50

	fmt.Printf("%v %T \n", n, n) // 30 int
	fmt.Printf("%v %T \n", a, a) // 10 int8
	fmt.Printf("%v %T \n", b, b) // 20 int16
	fmt.Printf("%v %T \n", c, c) // 40 int32
	fmt.Printf("%v %T \n", d, d) // 50 int64

	fmt.Printf("%v %T \n", w, w) // 10 uint8
	fmt.Printf("%v %T \n", x, x) // 20 uint16
	fmt.Printf("%v %T \n", y, y) // 40 uint32
	fmt.Printf("%v %T \n", z, z) // 50 uint64

	fmt.Printf("%v \n", n+int(b)) // 30
}
