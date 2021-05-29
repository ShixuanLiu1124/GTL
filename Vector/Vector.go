package Vector

import "GTL/Container"

type Vector interface {
	PushBack(interface{}) error

	PopBack() (interface{}, error)

	Set(int, int) error

	At(int) (interface{}, error)

	Remove(int, int) error

	Find(interface{}, func(interface{}, interface{}) bool) int

	Container.Container
}
