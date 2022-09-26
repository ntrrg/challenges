package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := bufio.NewScanner(os.Stdin)

	s.Scan()
	title := s.Text()
	s.Scan()
	author := s.Text()
	s.Scan()
	price, _ := strconv.Atoi(s.Text())

	book := NewMyBook(title, author, price)
	book.display()
}

type Book struct {
	title, author string
}

func (b Book) display() {
	fmt.Println("Implement the 'display' method!")
}

type MyBook struct {
	Book
	price int
}

func NewMyBook(title, author string, price int) MyBook {
	return MyBook{Book{title, author}, price}
}

func (b *MyBook) display() {
	fmt.Printf("Title: %v\nAuthor: %v\nPrice: %v\n", b.title, b.author, b.price)
}
