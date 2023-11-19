package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)


func cube(f float64) float64{
	return math.Pow(f, 3)
}

func f1(u uint) (uint, uint) {
	var factorial uint = 1
	 var sum uint = 0

	for i:=u;i>0;i-- {
		factorial *= i
	}

	for i:= 0;i<=int(u);i++ {
		sum += uint(i)
	}

	return factorial, sum
}

func myFunc(s string) int {
	a, b, c := s, s+s, s+s+s
	
	ai, err := strconv.Atoi(a)
	if err != nil {
		log.Fatal(err)
	}
	bi, err := strconv.Atoi(b)
	if err != nil {
		log.Fatal(err)
	}
	ci, err := strconv.Atoi(c)
	if err != nil {
		log.Fatal(err)
	}

	return ai+bi+ci
}


func sum(i ...int) (ret int) {
	for _, v := range i {
		ret+=v
	}

	return
}

func searchItem(a []string, b string) bool {
	for _, v :=range a {
		if strings.EqualFold(v, b) {
			return true
		}
	}
	return false
}

func print(msg string) {
	fmt.Println(msg)
}

func f2(i int) {
	fmt.Println(i)
}

 
func main() {
	c := cube(3.)
	fmt.Printf("Ex1: %f \n", c)

	f, s := f1(uint(5))
	fmt.Printf("Ex2: factorial: %d, Sum: %d \n", f, s)

	strSum := myFunc("5")
	fmt.Printf("Ex3: %d \n", strSum)

	ss := sum(1,2,3,4,5)
	fmt.Printf("Ex4: %d \n", ss)

	animals := []string{"Lion", "tiger", "bear"}
	result := searchItem(animals, "beaR")
	fmt.Println(result) // => true
	result = searchItem(animals, "lion")
	fmt.Println(result)     // => true

	defer print("The Go gopher is the iconic mascot of the Go project.")
	fmt.Println("Hello, Go playground!")

	x := f2
	fmt.Printf("%T\n", x) // => func(int)
	x(8)


}