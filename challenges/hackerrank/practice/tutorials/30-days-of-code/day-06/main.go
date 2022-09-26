package main

import "fmt"

func main() {
	var T int
	fmt.Scan(&T)

	for i := 1; i <= T; i++ {
		var S string
		fmt.Scan(&S)

		l := len(S) / 2
		even := make([]rune, 0, l)
		odd := make([]rune, 0, l+1)

		for j, c := range S {
			if j%2 == 0 {
				even = append(even, c)
			} else {
				odd = append(odd, c)
			}
		}

		fmt.Printf("%v %v\n", string(even), string(odd))
	}
}
