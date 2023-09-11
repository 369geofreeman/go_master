package main

import "fmt"

func main() {
	ex1()
	ex2()
	ex3()
	ex4()
	ex5()
	ex6()
}

func ex1() {
	for i := 0; i < 50; i++ {
		if i%7 == 0 {
			fmt.Println(i)
		}
	}
}

func ex2() {
	i := 0
	for {
		if i%7 == 0 {
			fmt.Println(i)
		}
		if i == 50 {
			break
		}
		i++
	}
}

func ex3() {
	limit := 0
	for i := 0; i < 50; i++ {
		if i%7 == 0 {
			fmt.Println(i)
			limit++
		}
		if limit == 3 {
			break
		}
	}
}

func ex4() {
	for i := 1; i < 500; i++ {
		if i%7 == 0 && i%5 == 0 {
			fmt.Println(i)
		}
	}
}

func ex5() {
	birthYear, currentYear := 1980, 2020

	for i := birthYear; i <= currentYear; {
		fmt.Printf("%d ", i)
		i++
	}
	fmt.Println()
}

func ex6() {
	age := -9
	switch {
	case age < 0, age > 100:
		fmt.Println("Invalid Age")
	case age < 18:
		fmt.Println("You are minor!")
	case age == 18:
		fmt.Println("Congratulations! You've just become major!")
	default:
		fmt.Println("You are major!")
	}
}
