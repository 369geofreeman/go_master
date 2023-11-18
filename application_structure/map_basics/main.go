package main

import "fmt"

func main() {
	var employees map[string]string
	fmt.Printf("%#v\n", employees)

	fmt.Printf("No of employees:%s\n", len(employees))

	fmt.Printf("The value for key %q is %q \n", "Dan", employees["Dan"])

	var accounts map[string]float64

	fmt.Printf("%#v\n", accounts["fg"])

	var m1 map[[5]int]string
	_ = m1

	// employees["Dan"] = "programmer"

	people := map[string]float64{}

	people["John"] = 21.4
	people["Mary"] = 10

	fmt.Println(people)

	map1 := make(map[int]int)
	map1[4] = 8

	balances := map[string]float64{
		"USD": 2134.213,
		"EUR": 37567.98,
	}

	fmt.Println(balances)

	m := map[int]int{1:10,2:20}
	_ = m


	balances["USD"] = 12.12
	balances["GBP"] = 100
	fmt.Println(balances)

	fmt.Println(balances["RON"])

	balances["RON"] = 0

	v, ok := balances["RON"]
	if ok {
		fmt.Println("RON HERE", v)
	} else {
		fmt.Println(v)
	}

	for k, v := range balances {
		fmt.Printf("%v, %v\n", k, v)
	}

	delete(balances, "USD")

	fmt.Println(balances)


}