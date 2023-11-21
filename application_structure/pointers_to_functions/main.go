package main

import "fmt"


func change(a *int) *float64{
	*a = 100

	b := 5.5
	return &b
}

func changeVar(i int) {
	i = 66
}


func main() {
	x := 8
	p := &x

	fmt.Println("Value of x before challing change():", x)
	change(p)
	fmt.Println("Value of x after challing change():", x)

	fmt.Println("Value of x before challing changeVar():", x)
	changeVar(x)
	fmt.Println("Value of x after challing change():", x)

}