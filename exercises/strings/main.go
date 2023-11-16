package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	ex1()
	ex2()
	ex3()
	ex4()
	ex5()
	ex6()
}

func ex1() {
	var name string = "bob"
	country := "Ireland"
	fmt.Printf("Your name: %s\nCountry: %s\n", name, country)
	fmt.Println("He says \"Hello\"")
	fmt.Println("C:\\Users\\a.txt")
}

func ex2() {
	r := 'ă'
	fmt.Println(r)

	s1, s2 := "ma", "m"

	ret := s1 + s2+ string(r)
	fmt.Println(ret)
}

func ex3() {
	s1 := "țară means country in Romanian"

	for i:=0;i<len(s1);i++ {
		fmt.Printf("%v | %T \n", s1[i], s1[i])
	}

	for _, i := range s1 {
		fmt.Printf("%v | %T \n", i, i)
	}
}

func ex4() {
	s1 := "Go is cool!"
	x := s1[0]
	fmt.Println(x)
 
	// s1[0] = 'x'
 
	// printing the number of runes in the string
	fmt.Println(utf8.RuneCountInString(s1))
}

func ex5() {
	s := "你好 Go!"
	b := []byte(s)
	fmt.Println(b)

	for i, v := range b {
		fmt.Printf("Index: %d, byte: %d \n", i, v)
	}
}

func ex6() {
	s := "你好 Go!"
	r := []rune(s)
	fmt.Println(r)

	for i, v := range r {
		fmt.Printf("Index: %d, rune: %d \n", i, v)
	}
}