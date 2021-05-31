package PriorityQueue

import (
	"GTL/Container"
)

type PriorityQueue interface {
	Push(values interface{}) error

	Pop() (interface{}, error)

	Top() (interface{}, error)

	swap(i, j int)

	up(index int)

	down(start int, end int) bool

	fix(index int)

	SetFunc(less func(interface{}, interface{}) bool)

	Container.Container
}
