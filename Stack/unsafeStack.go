/*
 *  Copyright (C) 2021  Shixuan Liu
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU General Public License for more details.
 *
 *     You should have received a copy of the GNU General Public License
 *     along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

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
	rear    *sNode
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
		rear:    node,
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
		rear:    node,
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
		prev:  s.rear,
	}
	s.rear.next = node
	s.rear = node
	s.size++

	return nil
}

func (s *unsafeStack) Top() (interface{}, error) {
	if s.Empty() {
		return nil, errors.New("This stack is empty")
	}

	return s.rear.value, nil
}

func (s *unsafeStack) Pop() (interface{}, error) {
	if s.Empty() {
		return nil, errors.New("This stack is empty")
	}

	value := s.rear.value
	s.rear = s.rear.prev
	s.rear.next = nil
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

// CatFromSlice 从slice中复制元素到Stack后面
func (s *unsafeStack) CatFromSlice(values []interface{}) error {
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
	s.rear = s.head
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

	// 使用 UseNumber 方法后，json包会将数字转换成一个内置的 Number 类型（而不是 float64），
	// 这个 Number 类型提供了转换为 int64、float64 等多个方法。
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
