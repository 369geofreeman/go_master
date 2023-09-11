package main

import "fmt"

func main() {
	ex1()
	ex2()
	ex3()
}

func ex1() {
	x, y, z := 10, 15.5, "Gophers"
	score := []int{10, 20, 30}

	fmt.Printf("%d\n", x)
	fmt.Printf("%f\n", y)
	fmt.Printf("%s\n", z)
	fmt.Printf("%#v\n", score)
	fmt.Printf("%v\n", score)

	fmt.Println(z)

	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%v\n", z)
	fmt.Printf("%v\n", score)

	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", score)
}

func ex2() {
	const x float64 = 1.422349587101
	fmt.Printf("%.4f", x)
}

func ex3() {
	shape := "circle"
	radius := 3.2
	const pi float64 = 3.14159
	circumference := 2 * pi * radius
 
	fmt.Printf("Shape: %s\n", shape)
	fmt.Printf("Circle's circumference with radius %f is %f\n", radius, circumference)
	_ = shape
}