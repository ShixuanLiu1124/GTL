package PriorityQueue

import (
	"GTL/Container"
)

type PriorityQueue interface {
	Push(interface{}) error

	Pop() (interface{}, error)

	Top() (interface{}, error)

	swap(int, int)

	up(int)

	down(int, int) bool

	fix(int)

	SetFunc(func(interface{}, interface{}) bool)

	Container.Container
}
