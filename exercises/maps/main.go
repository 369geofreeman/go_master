package main

import "fmt"

func main() {
	ex1()
	ex2()
	ex3()
}


func ex1() {
	var m1 map[string]int
	fmt.Printf("Type: %T, Value %v \n", m1, m1)

	m2 := map[int]string{1: "A", 2: "B"}

	m2[10] = "ABBA"

	fmt.Printf("Existing key: %s \n", m2[10])
	fmt.Printf("Non-existing key: %s \n", m2[11])
}

func ex2() {
	var m1 = map[int]bool{}
	m1[5] = true
	// _ = m1
 
	m2 := map[int]int{3: 10, 4: 40}
	m3 := map[int]int{3: 10, 4: 40}

	sm2 := fmt.Sprintf("%v", m2)
	sm3 := fmt.Sprintf("%v", m3)
 
	fmt.Println(sm2 == sm3)
}

func ex3() {
	m := map[int]bool{1: true, 2: false, 3: false}

	delete(m, 2)

	for k, v := range m {
		fmt.Printf("Key: %d, Value: %v \n", k, v)
	}
}