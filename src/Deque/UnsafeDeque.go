package Deque

import (
	"errors"
)

type DQNode struct {
	value interface{}
	next  *DQNode
	prev  *DQNode
}

type UnsafeDeque struct {
	size    int
	maxSize int
	head    *DQNode
	rail    *DQNode
}

func New(maxSize int, values ...interface{}) (*UnsafeDeque, error) {
	if maxSize != -1 && len(values) > maxSize {
		return nil, errors.New("Length of values is too long.")
	}

	node := &DQNode{
		value: nil,
		next:  nil,
		prev:  nil,
	}

	q := &UnsafeDeque{
		size:    0,
		maxSize: maxSize,
		head:    node,
		rail:    node,
	}

	for _, value := range values {
		q.PushBack(value)
	}

	return q, nil
}

func (q *UnsafeDeque) PushFront(value interface{}) error {
	if q.Fill() {
		return errors.New("This deque is fill.")
	}

	node := &DQNode{
		value: value,
		next:  nil,
		prev:  q.head,
	}
	node.next = q.head.next
	q.head.next = node
	q.size++

	return nil
}

func (q *UnsafeDeque) PushBack(value interface{}) error {
	if q.Fill() {
		return errors.New("This deque is fill.")
	}

	node := &DQNode{
		value: value,
		next:  nil,
		prev:  q.rail,
	}
	q.rail.next = node
	q.rail = node
	q.size++

	return nil
}

func (q *UnsafeDeque) Front() (interface{}, error) {
	if q.Empty() {
		return nil, errors.New("This deque is empty.")
	}

	return q.head.next.value, nil
}

func (q *UnsafeDeque) Back() (interface{}, error) {
	if q.Empty() {
		return nil, errors.New("This deque is empty.")
	}

	return q.rail.value, nil
}

func (q *UnsafeDeque) PopFront() (interface{}, error) {
	if q.Empty() {
		return nil, errors.New("This deque is empty")
	}

	value := q.head.next.value
	q.head.next = q.head.next.next
	q.size--

	return value, nil
}

func (q *UnsafeDeque) PopBack() (interface{}, error) {
	if q.Empty() {
		return nil, errors.New("This deque is empty")
	}

	value := q.rail.value
	q.rail = q.rail.prev
	q.size--

	return value, nil
}

/*---------------------------------以下为接口实现---------------------------------------*/

func (q *UnsafeDeque) CopyFromArray(values []interface{}) error {
	l := len(values)
	if q.maxSize != -1 && q.size+l > q.maxSize {
		return errors.New("Not enough free space.")
	}

	for a := range values {
		err := q.PushBack(a)
		if err != nil {
			return err
		}
	}
	q.size += l

	return nil
}

func (q *UnsafeDeque) Fill() bool {
	f := false
	if q.MaxSize() != -1 {
		f = q.Size() == q.MaxSize()
	}

	return f
}

func (q *UnsafeDeque) Empty() bool {
	return q.Size() == 0
}

func (q *UnsafeDeque) Size() int {
	return q.size
}

func (q *UnsafeDeque) MaxSize() int {
	return q.maxSize
}

func (q *UnsafeDeque) SetMaxSize(maxSize int) error {
	if maxSize != -1 && maxSize < q.size {
		return errors.New("New maxSize is less than current size.")
	}

	q.maxSize = maxSize

	return nil
}

func (q *UnsafeDeque) Clear() bool {
	q.rail = q.head
	q.head.prev = nil
	q.head.next = nil
	q.head.value = nil
	q.size = 0

	return true
}
