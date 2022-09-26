package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	fmt.Println(factorial(N))
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}

	return n * factorial(n-1)
}
