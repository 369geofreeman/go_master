package main

import "fmt"

type Article struct {
	Title  string
	Author string
}

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (a Article) String() string {
	return fmt.Sprintf("The %q acticle was written by %s.", a.Title, a.Author)
}

func (b Book) String() string {
	return fmt.Sprintf("The book %q was written by %s. it has %d pages.", b.Title, b.Author, b.Pages)
}

type Stringer interface {
	String() string
}

func main() {
	a := Article{
		Title:  "Understanding interfaces in GO",
		Author: "God",
	}

	Print(a)

	b := Book{
		Title:  "All About GO!",
		Author: "Jenny Dolph",
		Pages:  69,
	}

	Print(b)
}

func Print(s Stringer) {
	fmt.Println(s.String())
}
