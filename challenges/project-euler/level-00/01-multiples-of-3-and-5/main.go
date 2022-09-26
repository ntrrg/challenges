package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Scan(&n)
	n = n - 1
	fmt.Println(SummationBy(n, 3) + SummationBy(n, 5) - SummationBy(n, 15))
}

func Summation(n int) int {
	return n * (n + 1) / 2
}

func SummationBy(n, m int) int {
	return Summation(n/m) * m
}
