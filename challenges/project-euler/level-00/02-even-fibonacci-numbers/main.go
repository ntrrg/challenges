package main

import (
	"fmt"
)

func main() {
	var (
		n, r int
	)

	fmt.Scan(&n)

	n1, n2, n3 := 1, 1, 2
	for n3 < n {
		r += n3
		n1 = n2 + n3
		n2 = n1 + n3
		n3 = n1 + n2
	}

	fmt.Println(r)
}
