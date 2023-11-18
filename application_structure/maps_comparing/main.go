package main

import "fmt"

func main() {
	a := map[string]string{"A": "x"}
	b := map[string]string{"B": "y"}

	s1 := fmt.Sprintf("%v", a)
	s2 := fmt.Sprintf("%v", b)

	fmt.Println(s1)
	fmt.Println(s2)

	if s1 == s2 {
		fmt.Println("Maps are equal")
	} else {
		fmt.Println("Maps are not equal")
	}



}