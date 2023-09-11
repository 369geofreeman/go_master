package main

import (
	"fmt"
	"strconv"
)

func main() {
	ex1()
	ex2()
	ex3()
	ex4()
	ex5()
}

func ex1() {
	var i int = 3
	var f float64 = 3.2

	i2 := float64(i)
	f2 := int(f)

	fmt.Println(i2, f2)

}

func ex2() {
	var i = 3
	var f = 3.2
	var s1, s2 = "3.14", "5"

	i2 := fmt.Sprintf("%d", i)
	s2i, _ := strconv.Atoi(s2)
	f2 := fmt.Sprintf("%f", f)
	s1f, _ := strconv.ParseFloat(s1, 64)

	fmt.Println(i2, s2i, f2, s1f)
}

func ex3() {
	x, y := 4, 5.1

	z := float64(x) * y
	fmt.Println(z)

	a := 5.0
	b := 6.2 * a
	fmt.Println(b)
}

func ex4() {
	const (
		speed    float64 = 299792458
		distance float64 = 149600000
	)
	time := distance / speed
	fmt.Printf("%.6f\n", time)
}

func ex5() {
	x, y := 0.1, 5
    var z float64

    result1 := x < z || int(x) != int(z)
    fmt.Println(result1)
    
    result2 := y == 1 * 5 && int(z) == 0
    fmt.Println(result2)
}
