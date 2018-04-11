// This challenge is intended for object oriented languages, but this file aims
// to demonstrate that it is easy (and almost the same) to use the composition
// approach in Go.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)

	s.Scan()
	title := s.Text()
	s.Scan()
	author := s.Text()
	s.Scan()
	price := s.Text()

	book := NewMyBook(title, author, price)
	book.display()
}

type Book struct {
	title, author string
}

func NewBook(title, author string) *Book {
	fmt.Println("Do not attempt to directly instantiate an abstract class.")

	return nil
}

func (b *Book) display() {
	fmt.Println("Implement the 'display' method!")
}

type MyBook struct {
	Book
	price string
}

func NewMyBook(title, author, price string) *MyBook {
	b := new(MyBook)
	b.title = title
	b.author = author
	b.price = price

	return b
}

func (b *MyBook) display() {
	fmt.Printf("Title: %v\nAuthor: %v\nPrice: %v\n", b.title, b.author, b.price)
}
