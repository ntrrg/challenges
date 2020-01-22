package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)

	s.Scan()
	fmt.Println("Hello, World.")
	fmt.Println(s.Text())
}
