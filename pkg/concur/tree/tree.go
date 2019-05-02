package tree

import (
	"math/rand"
)

type Tree struct {
	Left  *Tree
	Right *Tree
	Value int
}

func Walk(t *Tree, ch chan int) {
	if t == nil {
		return
	}

	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

func New(n int) (root *Tree) {
	for i := range rand.Perm(n) {
		root = insert(root, i)
	}
	return
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, nil, v}
	}

	if t.Value > v {
		insert(t.Left, v)
	} else {
		insert(t.Right, v)
	}

	return t
}
