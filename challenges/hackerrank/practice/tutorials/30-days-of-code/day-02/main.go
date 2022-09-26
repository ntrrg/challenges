package main

import (
	"fmt"
)

func main() {
	var (
		mealCost float64

		tipPercent, taxPercent int

		total float64
	)

	fmt.Scan(&mealCost)
	total += mealCost

	fmt.Scan(&tipPercent)
	total += float64(tipPercent) / 100 * mealCost

	fmt.Scan(&taxPercent)
	total += float64(taxPercent) / 100 * mealCost

	fmt.Println(int64(total + 0.5))
}
