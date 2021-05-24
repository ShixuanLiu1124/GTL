package Stack

import (
	"errors"
)

type SNode struct {
	value interface{}
	next  *SNode
	prev  *SNode
}

type UnsafeStack struct {
	size    int
	maxSize int
	head    *SNode
	rail    *SNode
}

func New(maxSize int, values ...interface{}) (*UnsafeStack, error) {
	if maxSize != -1 && len(values) > maxSize {
		return nil, errors.New("Length of values is too long.")
	}

	node := &SNode{
		value: nil,
		next:  nil,
		prev:  nil,
	}

	s := &UnsafeStack{
		size:    0,
		maxSize: maxSize,
		head:    node,
		rail:    node,
	}

	for _, value := range values {
		s.Push(value)
	}

	return s, nil
}

func (s *UnsafeStack) Push(value interface{}) error {
	if s.Fill() {
		return errors.New("This stack is fill")
	}

	node := &SNode{
		value: value,
		next:  nil,
		prev:  s.rail,
	}
	s.rail.next = node
	s.rail = node
	s.size++

	return nil
}

func (s *UnsafeStack) Top() (interface{}, error) {
	if s.Empty() {
		return nil, errors.New("This stack is empty")
	}

	return s.rail.value, nil
}

func (s *UnsafeStack) Pop() (interface{}, error) {
	if s.Empty() {
		return nil, errors.New("This stack is empty")
	}

	value := s.rail.value
	s.rail = s.rail.prev
	s.rail.next = nil
	s.size--

	return value, nil
}

func (s *UnsafeStack) SetMaxSize(maxSize int) error {
	if maxSize != -1 && maxSize < s.size {
		return errors.New("New maxSize is less than current size.")
	}

	s.maxSize = maxSize

	return nil
}

func (s *UnsafeStack) CopyFromArray(values []interface{}) error {
	l := len(values)
	if s.maxSize != -1 && s.size+l > s.maxSize {
		return errors.New("Not enough free space.")
	}

	for a := range values {
		err := s.Push(a)
		if err != nil {
			return err
		}
	}
	s.size += l

	return nil
}

func (s *UnsafeStack) Fill() bool {
	f := false
	if s.maxSize != -1 && s.size == s.maxSize {
		f = true
	}
	return f
}

func (s *UnsafeStack) Empty() bool {
	return s.size == 0
}

func (s *UnsafeStack) Size() int {
	return s.size
}

func (s *UnsafeStack) MaxSize() int {
	return s.maxSize
}

func (s *UnsafeStack) ToString() string {
	// TODO ToString method

	return ""
}

func (s *UnsafeStack) Clear() bool {
	s.rail = s.head
	s.head.prev = nil
	s.head.next = nil
	s.head.value = nil
	s.size = 0

	return true
}
