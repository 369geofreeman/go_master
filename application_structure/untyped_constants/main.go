package main

import "fmt"

func main() {
	// typed constant
	const a float64 = 5.1

	//  untyped constant (we didn't specify it's type)
	const b = 6.7

	const c float64 = a * b
	const str = "hello" + "GO"

	const d = 5 > 10

	fmt.Println(d)

	// const x int = 5
	// const y float64 = 2.2 * x

	//  this works becaiuse both x and y are untyped constants
	const x = 5
	const y = 2.2 * 5
	fmt.Printf("%T\n", y)


	var i int = x  // x is typeless and changes to int
	var j float64 = x // x is typlesss and changes to float64. Behind teh scenes it does: var j float64 = float64(x)
	var p byte = x

	fmt.Println(i, j, p)

	// default types: If a type is needed but none is proviided
	const r = 5
	var rr = r 
	fmt.Printf("%T\n", rr)

}