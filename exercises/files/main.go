package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
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
	file, err := os.Create("info.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
}

func ex2() {
	oldpath, newpath := "info.txt", "data.txt"
	
	_, err := os.Stat(oldpath)
	if os.IsNotExist(err) {
		log.Fatal(err)
	}

	err = os.Rename(oldpath, newpath)
	if err != nil {
		log.Fatal(err)
	}

}

func ex3() {
	err := os.Remove("data.txt")
	if err != nil {
		log.Fatal(err)
	}
}

func ex4() {
	file, err := os.Open("a.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
    if err != nil {
        log.Fatal(err)
    }

	fmt.Println(string(data))
}

func ex5() {
	file, err := os.Open("a.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
        fmt.Println(scanner.Text())
    }

	if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

func ex6() {
	file, err := os.OpenFile(
        "info.txt",
        os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
        0644,
    )
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

	_, err = file.WriteString("The Go gopher is an iconic mascot!")
	if err != nil {
        log.Fatal(err)
    }
}