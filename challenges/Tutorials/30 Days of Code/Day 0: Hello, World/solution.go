package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, World.")

	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	fmt.Println(s.Text())
}
