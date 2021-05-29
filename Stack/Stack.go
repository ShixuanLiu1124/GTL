package Stack

import "GTL/Container"

type Stack interface {
	Push(value interface{}) error

	Top() (interface{}, error)

	Pop() (interface{}, error)

	Container.Container
}
