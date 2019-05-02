package tree

import (
	"math/rand"
)

type Tree struct {
	Left  *Tree
	Right *Tree
	Value int
}

func Same(a *Tree, b *Tree) bool {
	cha := make(chan int, 10)
	chb := make(chan int, 10)

	go func() {
		defer close(cha)
		Walk(a, cha)
	}()

	go func() {
		defer close(chb)
		Walk(b, chb)
	}()

	for {
		av, aok := <-cha
		bv, bok := <-chb

		if aok && bok {
			// both ok and having some value
			if av != bv {
				return false
			}
		} else if !(aok || bok) {
			// both are closed
			return true
		} else {
			// one is closed
			return false
		}
	}
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
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}

	return t
}
