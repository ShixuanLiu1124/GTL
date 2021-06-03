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

package PriorityQueue

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// unsafePriorityQueue 实现了一个小顶堆(根据less函数而定)
type unsafePriorityQueue struct {
	maxSize int
	s       []interface{}
	less    func(i, j interface{}) bool
}

func NewUnsafePriorityQueue(maxSize int, values ...interface{}) (*unsafePriorityQueue, error) {
	if maxSize != -1 && len(values) > maxSize {
		return nil, errors.New("Length of values is too long.")
	}

	q := &unsafePriorityQueue{
		maxSize: maxSize,
		s:       values,
		less:    nil,
	}
	q.fix(0)

	return q, nil
}

func NewUnsafePriorityQueueWithSlice(maxSize int, values []interface{}) (*unsafePriorityQueue, error) {
	if maxSize != -1 && len(values) > maxSize {
		return nil, errors.New("Length of values is too long.")
	}

	q := &unsafePriorityQueue{
		maxSize: maxSize,
		s:       []interface{}{},
		less:    nil,
	}

	for _, value := range values {
		q.s = append(q.s, value)
	}
	q.fix(0)

	return q, nil
}

func (q *unsafePriorityQueue) swap(i, j int) {
	q.s[i], q.s[j] = q.s[j], q.s[i]
}

// up 将元素向上调整
func (q *unsafePriorityQueue) up(index int) {
	for {
		// i是该元素的父亲结点
		i := (index - 1) / 2
		if i == index || !q.less(index, i) {
			break
		}
		q.swap(i, index)
		index = i
	}
}

// down 将元素向下调整
func (q *unsafePriorityQueue) down(start, end int) bool {
	i := start
	for {
		j1 := 2*i + 1
		// j1 < 0 after int overflow
		if j1 >= end || j1 < 0 {
			break
		}
		// 获取左孩子
		j := j1
		if j2 := j1 + 1; j2 < end && q.less(j2, j1) {
			// 获取右孩子
			j = j2 // = 2*i + 2
		}
		if !q.less(j, i) {
			break
		}
		q.swap(i, j)
		i = j
	}
	return i > start
}

// fix 调整堆
func (q *unsafePriorityQueue) fix(index int) {
	if !q.down(index, q.Size()) {
		q.up(index)
	}
}

func (q *unsafePriorityQueue) Push(value interface{}) error {
	if q.Fill() {
		return errors.New("This queue is fill.")
	}

	q.s = append(q.s, value)
	q.up(q.Size() - 1)

	return nil
}

// Pop 删除并返回最小元素（根据less函数）
func (q *unsafePriorityQueue) Pop() (interface{}, error) {
	if q.Empty() {
		return nil, errors.New("This queue is empty")
	}

	value := q.s[len(q.s)-1]
	n := q.Size() - 1
	q.swap(0, n)
	q.s = q.s[:q.Size()-1]

	q.down(0, n)

	return value, nil
}

func (q *unsafePriorityQueue) Top() (interface{}, error) {
	if q.Empty() {
		return nil, errors.New("This priority queue is empty")
	}

	return q.s[0], nil
}

func (q *unsafePriorityQueue) SetFunc(less func(interface{}, interface{}) bool) {
	q.less = less
	q.fix(0)
}

func (q *unsafePriorityQueue) Fill() bool {
	f := false
	if q.maxSize != -1 {
		f = len(q.s) == q.maxSize
	}

	return f
}

func (q *unsafePriorityQueue) Empty() bool {
	return q.Size() == 0
}

func (q *unsafePriorityQueue) Size() int {
	return len(q.s)
}

func (q *unsafePriorityQueue) MaxSize() int {
	return q.maxSize
}

func (q *unsafePriorityQueue) SetMaxSize(maxSize int) error {
	if maxSize != -1 && maxSize < len(q.s) {
		return errors.New("New maxSize is less than current size.")
	}

	q.maxSize = maxSize

	return nil
}

func (q *unsafePriorityQueue) Clear() {
	q.s = nil
}

func (q *unsafePriorityQueue) String() string {
	return fmt.Sprintf("%v", q.s)
}

func (q *unsafePriorityQueue) CatFromSlice(values []interface{}) error {
	l := len(values)
	if q.maxSize != -1 && q.Size()+l > q.maxSize {
		return errors.New("Not enough free space.")
	}

	for _, value := range values {
		err := q.Push(value)
		if err != nil {
			return err
		}
	}

	q.fix(0)

	return nil
}

func (q *unsafePriorityQueue) ToSlice() []interface{} {
	// // 切片直接指向存储空间，所以要复制到临时变量中再返回
	b := make([]interface{}, len(q.s))
	copy(b, q.s)

	return b
}

func (q *unsafePriorityQueue) MarshalJSON() ([]byte, error) {
	items := make([]string, 0, q.Size())

	for elem := range q.s {
		b, err := json.Marshal(elem)
		if err != nil {
			return nil, err
		}

		items = append(items, string(b))
	}

	return []byte(fmt.Sprintf("[%s]", strings.Join(items, ","))), nil
}

func (q *unsafePriorityQueue) UnmarshalJSON(b []byte) error {
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
			err = q.Push(t)
			if err != nil {
				fmt.Println(err)
				return err
			}
		}
	}

	q.fix(0)

	return nil
}
