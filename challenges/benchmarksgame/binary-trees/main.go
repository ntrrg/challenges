package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"sync"
)

func main() {
	if cpup := os.Getenv("CPUPROFILE"); cpup != "" {
		f, err := os.Create(cpup)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can't create CPU profile file: %v", err)
			os.Exit(1)
		}

		defer f.Close()

		if err := pprof.StartCPUProfile(f); err != nil {
			fmt.Fprintf(os.Stderr, "Can't profile CPU usage: %v", err)
			os.Exit(1)
		}

		defer pprof.StopCPUProfile()
	}

	min := 4
	var max int
	fmt.Scan(&max)

	if max < min+2 {
		max = min + 2
	}

	{
		st := NewTree(max + 1)
		fmt.Printf("stretch tree of depth %d\t check: %d\n", max+1, st.Check())
	}

	llt := NewTree(max)

	var wg sync.WaitGroup
	msgs := make([]string, max+1)

	for cd := min; cd <= max; cd += 2 {
		wg.Add(1)

		go func(cd int) {
			defer wg.Done()

			pool := NewPool(cd)
			n := 1 << (max - cd + min)
			i := 0
			sum := 0

			for ; i < n; i++ {
				t := NewTreeWithPool(cd, pool)
				sum += t.Check()
			}

			msgs[cd] = fmt.Sprintf("%d\t trees of depth %d\t check: %d", i, cd, sum)
		}(cd)
	}

	wg.Wait()

	for cd := min; cd <= max; cd += 2 {
		fmt.Println(msgs[cd])
	}

	fmt.Printf("long lived tree of depth %d\t check: %d\n", max, llt.Check())

	if memp := os.Getenv("MEMPROFILE"); memp != "" {
		f, err := os.Create(memp)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can't create memory profile file: %v", err)
			os.Exit(1)
		}

		defer f.Close()

		if err := pprof.WriteHeapProfile(f); err != nil {
			fmt.Fprintf(os.Stderr, "Can't profile memory usage: %v", err)
			os.Exit(1)
		}
	}
}

type Tree struct {
	left, right *Tree
}

func NewTree(d int) Tree {
	return NewTreeWithPool(d, NewPool(d))
}

func NewTreeWithPool(d int, pool []Tree) Tree {
	if d < 1 {
		return Tree{}
	}

	base := 1 << d
	total := base<<1 - 2
	pool = pool[:0]

	for i := 0; i < base; i++ {
		pool = append(pool, Tree{})
	}

	for i := 0; i < total-2; i += 2 {
		pool = append(pool, Tree{&pool[i], &pool[i+1]})
	}

	return Tree{&pool[total-2], &pool[total-1]}
}

func (t *Tree) Check() int {
	if t.left != nil {
		return t.left.Check() + t.right.Check() + 1
	}

	return 1
}

func NewPool(d int) []Tree {
	total := 1<<(d+1) - 2
	return make([]Tree, total)
}
