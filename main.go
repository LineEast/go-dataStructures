package main

import (
	"github.com/LineEast/go-dataStructures/lists/singly"
	"github.com/davecgh/go-spew/spew"
)

func Compare(a, b int) bool {
	return a > b
}

func main() {
	l := singly.NewOrderdList(Compare)
	for i := 5; i >= 0; i-- {
		l.Push(i)
	}

	spew.Dump(l.List)
}
