package main

import "fmt"

func main() {
	ex1()
	ex2()
	ex3()
}

func ex1() {
	type duration int
	var hour duration

	fmt.Printf("%T \n", hour)
	fmt.Printf("%v \n", hour)

	hour = 3600
	fmt.Printf("%v \n", hour)
}

func ex2() {
	type duration int

	var hour duration = 3600
    minute := 60
    fmt.Println(hour != duration(minute))
}

func ex3() {
	type mile float64
	type kilometer  float64

	const m2km = 1.609

	var mileBerlinToParis mile = 655.3
	var kmBerlinToParis kilometer

	kmBerlinToParis = kilometer(mileBerlinToParis) * m2km

	fmt.Println(kmBerlinToParis)
}