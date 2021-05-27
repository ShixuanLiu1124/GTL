package Deque

type Deque interface {
	PushFront(value interface{}) error

	PushBack(value interface{}) error

	Front() (interface{}, error)

	Back() (interface{}, error)

	PopFront() (interface{}, error)

	PopBack() (interface{}, error)
}
