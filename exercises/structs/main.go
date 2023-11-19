package main

import "fmt"

func main() {
	ex1()
}


func ex1() {
	type grades struct {
		grade, course string
	}

	type person struct {
		name string
		age int
		favColors []string
		grades grades
	}

	me := person{name: "John", age: 33, favColors: []string{"Blue", "Green"}, grades: grades{course: "lego", grade: "A+"}}
	you := person{name: "Paul", age: 90, favColors: []string{"orange", "Green"}, grades: grades{course: "Math", grade: "A+"}}

	fmt.Printf("%#v\n", me)
	fmt.Printf("%#v\n", you)

	me.name = "Andrei"

	you.favColors = append(you.favColors, "red")
	fmt.Printf("%v\n", you.favColors)

	for _, v := range you.favColors {
		fmt.Println(v)
	}

	me.grades.course = "Golang"
	me.grades.grade = "98"

}