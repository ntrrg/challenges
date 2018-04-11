package main

import "fmt"

func main() {
	var T int
	fmt.Scan(&T)

	for i := 1; i <= T; i++ {
		var S, even, odd string
		fmt.Scan(&S)

		for j, c := range S {
			if j%2 == 0 {
				even += string(c)
			} else {
				odd += string(c)
			}
		}

		fmt.Printf("%v %v\n", even, odd)
	}
}
