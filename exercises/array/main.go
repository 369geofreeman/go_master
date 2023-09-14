package main

import "fmt"

func main() {
	ex1()
	ex2()
	ex3()
}

func ex1() {
	var cities [2]string
	fmt.Println(cities)

	grades := [3]float64{1, 2}
	fmt.Println(grades)

	taskDone := [1]bool{}
	fmt.Println(taskDone)

	for i:=0;i<len(cities);i++ {
		fmt.Println(cities[i])
	}

	for _, v := range grades {
		fmt.Println(v)
	}
}

func ex2() {
	nums := [...]int{30, -1, -6, 90, -6}

	for _, v := range nums {
		if v %2==0 {
			fmt.Println(v)
		}
	}
}

func ex3() {
	myArray := [3]float64{1.2, 5.6}
 
    myArray[2] = 6
 
    a := 10.0
    myArray[0] = a
 
    myArray[2] = 10.10
 
    fmt.Println(myArray)
}