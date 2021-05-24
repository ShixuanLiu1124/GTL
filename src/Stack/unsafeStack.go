package Stack

import (
	"errors"
	"fmt"
	"strings"
)

type sNode struct {
	value interface{}
	next  *sNode
	prev  *sNode
}

type unsafeStack struct {
	size    int
	maxSize int
	head    *sNode
	rail    *sNode
}

func NewUnsafeStack(maxSize int, values ...interface{}) (*unsafeStack, error) {
	if maxSize != -1 && len(values) > maxSize {
		return nil, errors.New("Length of values is too long.")
	}

	node := &sNode{
		value: nil,
		next:  nil,
		prev:  nil,
	}

	s := &unsafeStack{
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

func (s *unsafeStack) Push(value interface{}) error {
	if s.Fill() {
		return errors.New("This stack is fill")
	}

	node := &sNode{
		value: value,
		next:  nil,
		prev:  s.rail,
	}
	s.rail.next = node
	s.rail = node
	s.size++

	return nil
}

func (s *unsafeStack) Top() (interface{}, error) {
	if s.Empty() {
		return nil, errors.New("This stack is empty")
	}

	return s.rail.value, nil
}

func (s *unsafeStack) Pop() (interface{}, error) {
	if s.Empty() {
		return nil, errors.New("This stack is empty")
	}

	value := s.rail.value
	s.rail = s.rail.prev
	s.rail.next = nil
	s.size--

	return value, nil
}

func (s *unsafeStack) SetMaxSize(maxSize int) error {
	if maxSize != -1 && maxSize < s.size {
		return errors.New("New maxSize is less than current size.")
	}

	s.maxSize = maxSize

	return nil
}

func (s *unsafeStack) CopyFromArray(values []interface{}) error {
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

func (s *unsafeStack) Fill() bool {
	f := false
	if s.maxSize != -1 && s.size == s.maxSize {
		f = true
	}
	return f
}

func (s *unsafeStack) Empty() bool {
	return s.size == 0
}

func (s *unsafeStack) Size() int {
	return s.size
}

func (s *unsafeStack) MaxSize() int {
	return s.maxSize
}

func (s *unsafeStack) ToString() string {
	// TODO ToString method

	return ""
}

func (s *unsafeStack) Clear() bool {
	s.rail = s.head
	s.head.prev = nil
	s.head.next = nil
	s.head.value = nil
	s.size = 0

	return true
}

func (s *unsafeStack) String() string {
	var b strings.Builder
	b.WriteString("unsafeQueue{")

	for p := s.head.next; p != nil; p = p.next {
		if p != s.head.next {
			b.WriteString(", ")
		}
		b.WriteString(fmt.Sprintf("%v", p.value))
	}
	b.WriteString("}")
	fmt.Println(b.String())

	return b.String()
}
