package PriorityQueue

import (
	"GTL/Container"
)

type PriorityQueue interface {
	Push(interface{}) error

	Pop() (interface{}, error)

	Top() (interface{}, error)

	Fix()

	SetFunc(func(i, j int) bool)

	Container.Container
}
