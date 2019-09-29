package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Scan(&n)
	n = n - 1
	fmt.Println(Summation(n/3)*3 + Summation(n/5)*5 - Summation(n/15)*15)
}

func Summation(n int) int {
	return n * (n + 1) / 2
}
