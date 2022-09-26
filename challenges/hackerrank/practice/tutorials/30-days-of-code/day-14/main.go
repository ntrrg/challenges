package main

import "fmt"

func main() {
	var N int

	fmt.Scan(&N)
	elements := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&elements[i])
	}

	d := NewDifference(elements)
	d.ComputeDifference()
	fmt.Println(d.MaximumDifference)
}

type Difference struct {
	MaximumDifference int

	elements []int
}

func NewDifference(a []int) *Difference {
	na := make([]int, len(a))
	copy(na, a)
	return &Difference{0, na}
}

func (d *Difference) ComputeDifference() {
	for i, l := 0, len(d.elements); i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			r := abs(d.elements[i] - d.elements[j])
			if r > d.MaximumDifference {
				d.MaximumDifference = r
			}
		}
	}
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}

	return n
}
