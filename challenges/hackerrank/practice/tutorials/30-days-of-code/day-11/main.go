package main

import "fmt"

func main() {
	A := [6][6]int{}

	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			fmt.Scan(&A[i][j])
		}
	}

	max := 0

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			hgSum := sum(A[i][j:j+3]...) + A[i+1][j+1] + sum(A[i+2][j:j+3]...)

			if hgSum > max || i == 0 && j == 0 {
				max = hgSum
			}
		}
	}

	fmt.Println(max)
}

func sum(nums ...int) int {
	var result int

	for _, v := range nums {
		result += v
	}

	return result
}
