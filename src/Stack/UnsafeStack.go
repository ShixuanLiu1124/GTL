package Stack

import (
	"errors"
)

type SNode struct {
	data interface{}
	next *SNode
	prev *SNode
}

type UnsafeStack struct {
	size    int
	maxSize int
	head    *SNode
	rail    *SNode
}

func New(maxSize int) *UnsafeStack {
	node := &SNode{
		data: nil,
		next: nil,
		prev: nil,
	}

	return &UnsafeStack{
		size:    0,
		maxSize: maxSize,
		head:    node,
		rail:    node,
	}
}

func (s *UnsafeStack) Push(data interface{}) error {
	if s.Fill() {
		return errors.New("This stack is fill")
	}

	node := &SNode{
		data: data,
		next: nil,
		prev: s.rail,
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

	return s.rail.data, nil
}

func (s *UnsafeStack) Pop() (interface{}, error) {
	if s.Empty() {
		return nil, errors.New("This stack is empty")
	}

	data := s.rail.data
	s.rail = s.rail.prev
	s.rail.next = nil
	s.size--

	return data, nil
}

func (s *UnsafeStack) SetMaxSize(maxSize int) error {
	if maxSize != -1 && maxSize < s.size {
		return errors.New("New maxSize is less than current size.")
	}

	s.maxSize = maxSize

	return nil
}

func (s *UnsafeStack) CopyFromArray(datas []interface{}) error {
	l := len(datas)
	if s.maxSize != -1 && s.size+l > s.maxSize {
		return errors.New("Not enough free space.")
	}

	for a := range datas {
		err := s.Push(a)
		if err != nil {
			return err
		}
	}
	s.size += l

	return nil
}

func (s *UnsafeStack) Copy(other *UnsafeStack) error {
	// TODO: Copy method

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
	s.head.data = nil
	s.size = 0

	return true
}
