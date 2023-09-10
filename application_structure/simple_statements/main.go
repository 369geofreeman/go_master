package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	i, err := strconv.Atoi("45a")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}

	if i, err := strconv.Atoi("x20"); err == nil {
		fmt.Println(i)
	} else {
		fmt.Println(err)
	}

	if args := os.Args; len(args) != 2 {
		fmt.Println("One argument is required")
	} else if km, err := strconv.Atoi(args[1]); err != nil {
		fmt.Println("Argument must be an intiger", err)
	} else {
		fmt.Printf("%d in miles is %f \n", km, float64(km)* 0.621)
	}
}