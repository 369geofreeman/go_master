package main

import (
	"fmt"
)

func main() {
	const days int = 7

	var i int

	fmt.Println(i)

	const pi float64 = 3.14
	const secondsInHour = 3600

	duration := 234 // in hours
	fmt.Printf("Duration in seconds: %v\n", duration*secondsInHour)

	x, y := 5, 1
	fmt.Println(x / y)

	const a = 5
	const b = 1

	fmt.Println(a / b)

	const n, m int = 4, 5
	const n1, m1 = 6, 7

	const (
		min1 = -500
		min2
		min3 
	)

	fmt.Println(min1, min2, min3)

	// constant rules:
	// 1. you can't change a constant 
	const temp int = 500
	// temp = 50

	// 2. you can not initiate a constant at runtime
	// const power = math.Pow(2, 3)

	// 3. You can not use a variable to initiate a constant
	// t := 5
	// const tc = t

	// 4. 
	const l1 = len("hello")

}
