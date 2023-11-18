package main

import "fmt"

func main() {
	friends := map[string]int{"Dan": 40, "Maria": 25}

	neighbours := friends

	friends["Dan"] = 50

	fmt.Println(neighbours)

	people := make(map[string]int)

	for k, v := range friends {
		people[k] = v
	}

	people["Anne"] = 18

	fmt.Printf("%#v ----------- %#v \n", people, friends)

	delete(friends, "Dan")

	fmt.Println(friends)

}