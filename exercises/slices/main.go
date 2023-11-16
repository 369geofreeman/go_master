package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	ex1()
	ex2()
	ex3()
	ex4()
	ex5()
	ex6()
	ex7()
	ex8()
}

func ex1() {
	s1 := []string{"1", "2", "3"}
	fmt.Println(s1)
}

func ex2() {
	mySlice := []float64{1.2, 5.6}
 
    mySlice[1] = 6
 
    a := float64(10)
    mySlice[0] = a
 
    mySlice[1] = 10.10
 
    mySlice = append(mySlice, a)
 
    fmt.Println(mySlice)
}

func ex3() {
	nums := []float64{1, 2, 3}
	nums = append(nums, 10.1)
	nums = append(nums, 4.1, 5.5, 6.6)
	fmt.Println(nums)

	n := []float64{9, 8}

	nums = append(nums, n...)
	fmt.Println(nums)
}

func ex4() {
	if len(os.Args) < 2 {
		log.Fatal("Please run the program with arguments (2-10 numbers)!")
		args := os.Args[1:]
		sum, product := 0., 1.

		if len(args) < 2 || len(args) > 10 {
			fmt.Println("Please enter between 2 and 10 numbers!")
		} else {
	
			for _, v := range args {
				num, err := strconv.ParseFloat(v, 64)
				if err != nil {
					continue //if it can't convert string to float64, continue with the next number
				}
				sum += num
				product *= num
			}
		}
	
		fmt.Printf("Sum: %v, Product: %v\n", sum, product)
	}
}

func ex5() {
	nums := []int{5, -1, 9, 10, 1100, 6, -1, 6}

	for _, v := range nums[1:len(nums)-1] {
		fmt.Println(v)
	}
}

func ex6() {
	friends := []string{"Marry", "John", "Paul", "Diana"}
	c := copy([]string{}, friends)

	friends[0] = "baboon"

	fmt.Printf("c: %v | friends: %v \n", c, friends)
}

func ex7() {
	friends := []string{"Marry", "John", "Paul", "Diana"}
	c := append(friends, friends...)

	friends[0] = "baboon"

	fmt.Printf("c: %v | friends: %v \n", c, friends)

}

func ex8() {
	years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}
	newYears := []int{}
	newYears = append(years[:3], years[len(years)-3:]...)
	fmt.Println(newYears)
}