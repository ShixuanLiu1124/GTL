package Deque

import (
	"errors"
	"fmt"
	"strings"
)

type dQNode struct {
	value interface{}
	next  *dQNode
	prev  *dQNode
}

type unsafeDeque struct {
	size    int
	maxSize int
	head    *dQNode
	rail    *dQNode
}

func New(maxSize int, values ...interface{}) (*unsafeDeque, error) {
	if maxSize != -1 && len(values) > maxSize {
		return nil, errors.New("Length of values is too long.")
	}

	node := &dQNode{
		value: nil,
		next:  nil,
		prev:  nil,
	}

	q := &unsafeDeque{
		size:    0,
		maxSize: maxSize,
		head:    node,
		rail:    node,
	}

	for _, value := range values {
		err := q.PushBack(value)
		if err != nil {
			return nil, err
		}
	}

	return q, nil
}

func (q *unsafeDeque) PushFront(value interface{}) error {
	if q.Fill() {
		return errors.New("This deque is fill.")
	}

	node := &dQNode{
		value: value,
		next:  nil,
		prev:  q.head,
	}
	node.next = q.head.next
	q.head.next = node
	q.size++

	return nil
}

func (q *unsafeDeque) PushBack(value interface{}) error {
	if q.Fill() {
		return errors.New("This deque is fill.")
	}

	node := &dQNode{
		value: value,
		next:  nil,
		prev:  q.rail,
	}
	q.rail.next = node
	q.rail = node
	q.size++

	return nil
}

func (q *unsafeDeque) Front() (interface{}, error) {
	if q.Empty() {
		return nil, errors.New("This deque is empty.")
	}

	return q.head.next.value, nil
}

func (q *unsafeDeque) Back() (interface{}, error) {
	if q.Empty() {
		return nil, errors.New("This deque is empty.")
	}

	return q.rail.value, nil
}

func (q *unsafeDeque) PopFront() (interface{}, error) {
	if q.Empty() {
		return nil, errors.New("This deque is empty")
	}

	value := q.head.next.value
	q.head.next = q.head.next.next
	q.size--

	return value, nil
}

func (q *unsafeDeque) PopBack() (interface{}, error) {
	if q.Empty() {
		return nil, errors.New("This deque is empty")
	}

	value := q.rail.value
	q.rail = q.rail.prev
	q.size--

	return value, nil
}

/*---------------------------------以下为接口实现---------------------------------------*/

func (q *unsafeDeque) CopyFromArray(values []interface{}) error {
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

func (q *unsafeDeque) Fill() bool {
	f := false
	if q.MaxSize() != -1 {
		f = q.Size() == q.MaxSize()
	}

	return f
}

func (q *unsafeDeque) Empty() bool {
	return q.Size() == 0
}

func (q *unsafeDeque) Size() int {
	return q.size
}

func (q *unsafeDeque) MaxSize() int {
	return q.maxSize
}

func (q *unsafeDeque) SetMaxSize(maxSize int) error {
	if maxSize != -1 && maxSize < q.size {
		return errors.New("New maxSize is less than current size.")
	}

	q.maxSize = maxSize

	return nil
}

func (q *unsafeDeque) Clear() bool {
	q.rail = q.head
	q.head.prev = nil
	q.head.next = nil
	q.head.value = nil
	q.size = 0

	return true
}

func (q *unsafeDeque) String() string {
	var b strings.Builder
	b.WriteString("unsafeQueue{")

	for p := q.head.next; p != nil; p = p.next {
		if p != q.head.next {
			b.WriteString(", ")
		}
		b.WriteString(fmt.Sprintf("%v", p.value))
	}
	b.WriteString("}")
	fmt.Println(b.String())

	return b.String()
}
