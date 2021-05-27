package PriorityQueue

import (
	"GTL/Container"
)

type PriorityQueue interface {
	Push(interface{}) error

	Pop() (interface{}, error)

	Top() (interface{}, error)

	Fix(int)

	SetFunc(func(interface{}, interface{}) bool)

	Container.Container
}
