package Stack

import (
	"bytes"
	"encoding/json"
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
		err := s.Push(value)
		if err != nil {
			return nil, err
		}
	}

	return s, nil
}

func NewUnsafeStackWithSlice(maxSize int, values []interface{}) (*unsafeStack, error) {
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
		err := s.Push(value)
		if err != nil {
			return nil, err
		}
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

/*---------------------------------以下为接口实现---------------------------------------*/

func (s *unsafeStack) SetMaxSize(maxSize int) error {
	if maxSize != -1 && maxSize < s.size {
		return errors.New("New maxSize is less than current size.")
	}

	s.maxSize = maxSize

	return nil
}

func (s *unsafeStack) CopyFromSlice(values []interface{}) error {
	l := len(values)
	if s.maxSize != -1 && s.size+l > s.maxSize {
		return errors.New("Not enough free space.")
	}

	fmt.Println("values =", values)

	for _, value := range values {

		fmt.Println("value =", value)

		err := s.Push(value)
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

func (s *unsafeStack) Clear() {
	s.rail = s.head
	s.head.prev = nil
	s.head.next = nil
	s.head.value = nil
	s.size = 0
}

func (s *unsafeStack) String() string {
	var b strings.Builder
	b.WriteString("unsafeStack{")

	for p := s.head.next; p != nil; p = p.next {
		if p != s.head.next {
			b.WriteString(", ")
		}
		b.WriteString(fmt.Sprintf("%v", p.value))
	}
	b.WriteString("}")

	return b.String()
}

// ToSlice 将切片以切片形式返回
func (s *unsafeStack) ToSlice() []interface{} {
	ans := make([]interface{}, s.size)

	for p := s.head.next; p != nil; p = p.next {
		ans = append(ans, p.value)
	}

	return ans
}

// MarshalJSON 将Stack中的所有元素以Json数组的形式返回
func (s *unsafeStack) MarshalJSON() ([]byte, error) {
	items := make([]string, 0, s.Size())

	for p := s.head.next; p != nil; p = p.next {
		b, err := json.Marshal(p.value)
		if err != nil {
			return nil, err
		}

		items = append(items, string(b))
	}

	return []byte(fmt.Sprintf("[%s]", strings.Join(items, ","))), nil
}

// UnmarshalJSON 从给定的Json数组中解析出一个Stack,数字将被解析为json.Number
func (s *unsafeStack) UnmarshalJSON(b []byte) error {
	var i []interface{}

	d := json.NewDecoder(bytes.NewReader(b))
	d.UseNumber()
	err := d.Decode(&i)
	if err != nil {
		return err
	}

	for _, value := range i {
		switch t := value.(type) {
		case []interface{}, map[string]interface{}:
			continue
		default:
			err = s.Push(t)
			if err != nil {
				fmt.Println(err)
				return err
			}
		}
	}

	return nil
}
