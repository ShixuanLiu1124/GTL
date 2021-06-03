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

package Deque

import (
	"bytes"
	"encoding/json"
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
	rear    *dQNode
}

func NewUnsafeDeque(maxSize int, values ...interface{}) (*unsafeDeque, error) {
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
		rear:    node,
	}

	for _, value := range values {
		err := q.PushBack(value)
		if err != nil {
			return nil, err
		}
	}

	return q, nil
}

func NewUnsafeDequeWithSlice(maxSize int, values []interface{}) (*unsafeDeque, error) {
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
		rear:    node,
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
		prev:  q.rear,
	}
	q.rear.next = node
	q.rear = node
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

	return q.rear.value, nil
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

	value := q.rear.value
	q.rear = q.rear.prev
	q.size--

	return value, nil
}

/*---------------------------------以下为接口实现---------------------------------------*/

// CatFromSlice 从slice中复制元素到Deque后面
func (q *unsafeDeque) CatFromSlice(values []interface{}) error {
	l := len(values)
	if q.maxSize != -1 && q.size+l > q.maxSize {
		return errors.New("Not enough free space.")
	}

	for _, value := range values {
		err := q.PushBack(value)
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

func (q *unsafeDeque) Clear() {
	q.rear = q.head
	q.head.prev = nil
	q.head.next = nil
	q.head.value = nil
	q.size = 0
}

func (q *unsafeDeque) String() string {
	var b strings.Builder
	b.WriteString("SafeQueue{")

	for p := q.head.next; p != nil; p = p.next {
		if p != q.head.next {
			b.WriteString(", ")
		}
		b.WriteString(fmt.Sprintf("%v", p.value))
	}
	b.WriteString("}")

	return b.String()
}

// ToSlice 将队列以切片形式返回
func (q *unsafeDeque) ToSlice() []interface{} {
	ans := make([]interface{}, q.size)

	for p := q.head.next; p != nil; p = p.next {
		ans = append(ans, p.value)
	}

	return ans
}

// MarshalJSON 将Deque中的所有元素以Json数组的形式返回
func (q *unsafeDeque) MarshalJSON() ([]byte, error) {
	items := make([]string, 0, q.Size())

	for p := q.head.next; p != nil; p = p.next {
		b, err := json.Marshal(p.value)
		if err != nil {
			return nil, err
		}

		items = append(items, string(b))
	}

	return []byte(fmt.Sprintf("[%s]", strings.Join(items, ","))), nil
}

// UnmarshalJSON 从给定的Json数组中解析出一个Deque,数字将被解析为json.Number
func (q *unsafeDeque) UnmarshalJSON(b []byte) error {
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
			err = q.PushBack(t)
			if err != nil {
				fmt.Println(err)
				return err
			}
		}
	}

	return nil
}
