package main

import (
	"fmt"
)

func main() {
	var n uint8
	fmt.Scan(&n)

	if n%2 != 0 || (n >= 6 && n <= 20) {
		fmt.Println("Weird")
	} else {
		fmt.Println("Not Weird")
	}
}
