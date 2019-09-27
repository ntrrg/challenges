package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		n int
		s string
	)
	fmt.Scan(&n)
	fmt.Scan(&s)

	l := findLongestPatternLen(&s)
	fmt.Println(l)
}

func findLongestPatternLen(s *string) int {
	if len(s) < 2 || (len(s) == 2 && s[0] == s[1]) {
		return 0
	}

	patterns := make(map[rune]*Pattern)
	ban := make(map[rune]bool)

	for _, v := range s {
		if p, ok := patterns[v]; !ok {

		}

		n := Node{v}
		fmt.Println(n)
	}
}

type Pattern struct {
	count  int
	values []*string
}
