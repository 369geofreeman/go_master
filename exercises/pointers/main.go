package main

import "fmt"

func main() {
	ex1()
	ex2()
	ex3()
}

func ex1() {
	x := 10.10
	fmt.Println(&x)

	ptr := &x
	fmt.Printf("Type: %T, value %p \n", ptr, ptr)

	fmt.Printf("Address: %v, value %v \n", &ptr, *ptr)
}

func ex2() {
	x, y := 10, 2
	ptrx, ptry := &x, &y

	z := *ptrx / *ptry

	fmt.Println(z)
}

func swap (x, y *float64) () {
	*x, *y = *y, *x
}

func ex3() {
	x, y := 5.5, 8.8

	fmt.Println("BEFORE =>", x, y)
	swap(&x, &y)
	fmt.Println("AFTER =>", x, y)
}
