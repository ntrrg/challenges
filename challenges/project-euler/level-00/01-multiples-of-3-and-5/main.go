package main

import (
	"fmt"
)

func main() {
	var n, r int

	fmt.Scanf("%d", &n)

	for n > 3 {
		n = n - 1

		if n%3 == 0 || n%5 == 0 {
			r += n
		}
	}

	fmt.Println(r)
}
