package Deque

import "GTL/Container"

type Deque interface {
	Container.Container

	PushFront(values interface{}) error

	PushBack(values interface{}) error

	Front() (interface{}, error)

	Back() (interface{}, error)

	PopFront() (interface{}, error)

	PopBack() (interface{}, error)
}
