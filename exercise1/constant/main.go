package main

import (
	"fmt"
	"math"
)

func main() {
	ex1()
	ex2()
	ex3()
	ex4()
}

func ex1() {
	const (
		daysInWeek = 7
		lightSpeed = 299792458
		pi         = 3.14159
		secPerDay = 86400
		daysYear = 365
	)

	fmt.Printf("Seconds in a year: %d \n", secPerDay * daysYear)
}

func ex2() {
	const x int = 10
 
    // declaring a constant of type slice int ([]int)
    var m = []int{1, 3, 4, 5, 6, 8}
    _ = m
}

func ex3() {
	const a int = 7
	const b float64 = 5.6
	const c = float64(a) * b
 
	x := 8
	var xc int = x
	_ = xc

	// var (
	// 	s1 float64 = 8
	// 	s2 float64 = 128
	// )
 
	var noIPv6 = math.Pow(8, 128)

	fmt.Println(noIPv6)
}

func ex4() {
	const (
		jun = iota + 6
		july
		august
	)

	fmt.Println(jun, july, august)
}