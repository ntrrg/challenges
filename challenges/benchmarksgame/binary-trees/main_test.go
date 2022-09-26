package main

import (
	"testing"
)

func TestTree(t *testing.T) {
	bt := NewTree(10)

	if c := bt.Check(); c != 2047 {
		t.Errorf("Invalid tree genarated, %v", c)
	}
}

const DEPTH = 10

func BenchmarkTree(b *testing.B) {
	pool := NewPool(DEPTH)

	for i := 0; i < b.N; i++ {
		x := NewTreeWithPool(DEPTH, pool)
		x.Check()
	}
}

func BenchmarkNewTree(b *testing.B) {
	pool := NewPool(DEPTH)

	for i := 0; i < b.N; i++ {
		NewTreeWithPool(DEPTH, pool)
	}
}

func BenchmarkTree_Check(b *testing.B) {
	pool := NewPool(DEPTH)

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x := NewTreeWithPool(DEPTH, pool)
		b.StartTimer()

		x.Check()
	}
}

func BenchmarkNode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := create(DEPTH)
		x.check()
	}
}

type node struct {
	next *next
}
type next struct {
	left, right node
}

func create(d int) node {
	if d == 1 {
		return node{&next{node{}, node{}}}
	}
	return node{&next{create(d - 1), create(d - 1)}}
}

func (p node) check() int {
	sum := 1
	current := p.next
	for current != nil {
		sum += current.right.check() + 1
		current = current.left.next
	}
	return sum
}
