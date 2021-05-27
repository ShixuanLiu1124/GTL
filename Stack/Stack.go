package Stack

import "GTL/Container"

type Stack interface {
	Container.Container

	Push(value interface{}) error

	Top() (interface{}, error)

	Pop() (interface{}, error)
}
