package main

import "fmt"

func main() {
	var n, count, max int
	fmt.Scan(&n)

	for _, v := range fmt.Sprintf("%b", n) {
		if v == '1' {
			count++

			if count > max {
				max = count
			}
		} else {
			count = 0
		}
	}

	fmt.Println(max)
}
