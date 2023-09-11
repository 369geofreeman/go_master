package main

import "fmt"

var version = "3.1"

func main() {
	sayHello()
	types()
	errorCheck()
}
func sayHello() {
	fmt.Println("Hello package v1.0.0!")

	var x int = 10
	_ = x
}

func types() {
	// var a int
	// var b float64
	// var c bool
	// var d string
	var (
		a int
		b float64
		c bool
		d string
	)

	// x := -20
	// y := 15.5
	// z := "Gopher!"
	x, y, z := -20, 15.5, "Gopher!"

	fmt.Println(a, b, c, d, x, z, y)
}

func errorCheck() {
	var a float64 = 7.1

	x, y := true, 3.7

	a, x = 5.5, false

	_, _, _ = a, x, y

	name := "Golang"
	fmt.Println(name)
}
