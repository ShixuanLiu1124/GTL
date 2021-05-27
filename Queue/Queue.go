package Queue

import "GTL/Container"

type Queue interface {
	Container.Container

	Push(value interface{}) error

	Front() (interface{}, error)

	Pop() (interface{}, error)
}
