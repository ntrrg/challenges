package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	A := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}

	for ; N > 1; N-- {
		fmt.Printf("%v ", A[N-1])
	}

	fmt.Printf("%v", A[0])
}
