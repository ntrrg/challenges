package main

import "fmt"

func main() {
	var (
		l int
		s string

		tc = NewTC()
	)

	fmt.Scan(&l)
	fmt.Scan(&s)

	for i := 0; i < l; i++ {
		tc.Add(s[i])
	}

	fmt.Println(tc.Longest())
}

type TC struct {
	letters map[byte]int
	pairs   map[string]int
}

func NewTC() *TC {
	return &TC{
		letters: make(map[byte]int),
		pairs:   make(map[string]int),
	}
}

func (tc *TC) Add(c byte) {
	if _, ok := tc.letters[c]; !ok {
		tc.addNew(c)
		return
	}

	tc.letters[c] += 1

	for p, n := range tc.pairs {
		if (p[0] != c && p[1] != c) || n == 0 {
			continue
		}

		n++

		if (p[0] == c && n%2 == 0) || (p[1] == c && n%2 == 1) {
			n = 0
		}

		tc.pairs[p] = n
	}
}

func (tc *TC) Longest() int {
	max := 0

	for _, n := range tc.pairs {
		if n > max {
			max = n
		}
	}

	return max
}

func (tc *TC) addNew(c byte) {
	for l, n := range tc.letters {
		if n != 1 {
			continue
		}

		p := string([]byte{l, c})
		tc.pairs[p] = 2
	}

	tc.letters[c] = 1
}
