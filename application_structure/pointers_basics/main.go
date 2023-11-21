package main

import "fmt"

func main() {
	name := "Andrei"
	fmt.Println(&name)

	var x int = 2
	ptr := &x

	fmt.Printf("Ptr is of type %T with a value of %v and address %p\n", ptr, ptr, &ptr)
	fmt.Printf("Address of ptr is %p\n", &x)
	
	var ptr1 *float64
	_ = ptr1

	p := new(int)
	x = 100
	p = &x

	fmt.Printf("P is of type %T with a value of %p \n", p, p)
	fmt.Printf("Address of p is %p\n", &x)

	*p = 90 // equlivant to x
	fmt.Println(x, *p)
	fmt.Println(*p == x)

	*p = 10 // x = 10
	*p = *p/2 // x/2
	fmt.Println(x)

	// &value => pointer
	// *pointer => value


}