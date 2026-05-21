package main

import "fmt"

func main() {
	var a int
	var b float64
	var c string
	var d bool

	// Default zero values for the variables
	fmt.Printf("a: %d\n", a) //0
	fmt.Printf("b: %f\n", b) //0.000000
	fmt.Printf("c: %q\n", c) //empty string ""
	fmt.Printf("d: %t\n", d) //false
}
