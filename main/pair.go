package main

import "fmt"

type Book struct {
}

func (this *Book) ReadBook() {
	fmt.Println("Reading this book.")
}

func main() {
	book := &Book{}
	fmt.Println(book)
	book.ReadBook()

	f := 1.234
	fmt.Printf("%T", f)
}
