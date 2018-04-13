package main

import (
	"fmt"
)

func main() {
	var mealCost, tip, tax float64
	fmt.Scan(&mealCost, &tip, &tax)

	tip /= 100
	tax /= 100
	total := mealCost + mealCost*tip + mealCost*tax

	fmt.Printf("The total meal cost is %v dollars.\n", round(total))
}

func round(n float64) int {
	return int(n + 0.5)
}
