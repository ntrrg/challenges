package main

import "fmt"

func main() {
	A := [][]int{}

	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			var n int
			fmt.Scan(&n)

			if j == 0 {
				A = append(A, []int{n})
			} else {
				A[i] = append(A[i], n)
			}
		}
	}

	max := 0

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			hg := []int{}
			hg = append(hg, A[i][j:j+3]...)
			hg = append(hg, A[i+1][j+1])
			hg = append(hg, A[i+2][j:j+3]...)

			hgSum := sum(hg...)

			if hgSum > max || i == 0 && j == 0 {
				max = hgSum
			}
		}
	}

	fmt.Println(max)
}

func sum(nums ...int) (result int) {
	for _, v := range nums {
		result += v
	}

	return result
}
