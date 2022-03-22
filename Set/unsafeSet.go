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

package Set

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// OrderedPair 表示一个二元组，用于求笛卡尔积
type OrderedPair struct {
	First  interface{}
	Second interface{}
}

type unsafeSet struct {
	m       map[interface{}]struct{}
	maxSize int
}

func NewUnsafeSet(maxSize int, values ...interface{}) (*unsafeSet, error) {
	if maxSize != -1 && len(values) > maxSize {
		return nil, errors.New("Length of values is too long.")
	}

	var mm map[interface{}]struct{}
	if maxSize != -1 {
		mm = make(map[interface{}]struct{}, maxSize)
	} else {
		mm = make(map[interface{}]struct{})
	}
	for _, v := range values {
		mm[v] = struct{}{}
	}

	return &unsafeSet{
		m:       mm,
		maxSize: maxSize,
	}, nil
}

func NewUnsafeSetWithSlice(maxSize int, values []interface{}) (*unsafeSet, error) {
	if maxSize != -1 && len(values) > maxSize {
		return nil, errors.New("Length of values is too long.")
	}

	var mm map[interface{}]struct{}
	if maxSize != -1 {
		mm = make(map[interface{}]struct{}, maxSize)
	} else {
		mm = make(map[interface{}]struct{})
	}
	for _, v := range values {
		mm[v] = struct{}{}
	}

	return &unsafeSet{
		m:       mm,
		maxSize: maxSize,
	}, nil
}

// Equal 用来判定两个OrderedPair对象是否相等
func (pair *OrderedPair) Equal(other OrderedPair) bool {
	if pair.First == other.First &&
		pair.Second == other.Second {
		return true
	}

	return false
}

// Insert 向集合中添加元素
func (s *unsafeSet) Insert(value interface{}) error {
	if s.Fill() {
		return errors.New("This set is fill.")
	}

	s.m[value] = struct{}{}

	return nil
}

func (s *unsafeSet) Contains(values ...interface{}) bool {
	for _, val := range values {
		if _, ok := s.m[val]; !ok {
			return false
		}
	}
	return true
}

// IsSubset 判断other是否是s的子集
func (s *unsafeSet) IsSubset(other Set) bool {
	_ = other.(*unsafeSet)
	if s.Size() > other.Size() {
		return false
	}
	for elem := range s.m {
		if !other.Contains(elem) {
			return false
		}
	}
	return true
}

// IsProperSubset 判断other是否是s的真子集
func (s *unsafeSet) IsProperSubset(other Set) bool {
	return s.IsSubset(other) && !s.Equal(other)
}

// IsSuperset 判断other是否是s的超集
func (s *unsafeSet) IsSuperset(other Set) bool {
	return other.IsSubset(s)
}

// IsProperSuperset 判断other是否是s的真超集
func (s *unsafeSet) IsProperSuperset(other Set) bool {
	return s.IsSuperset(other) && !s.Equal(other)
}

// Union 求该集合s和other的并集
func (s *unsafeSet) Union(other Set) Set {
	o := other.(*unsafeSet)

	unionedSet, _ := NewUnsafeSet(s.MaxSize() + other.MaxSize())

	for elem := range s.m {
		_ = unionedSet.Insert(elem)
	}
	for elem := range o.m {
		_ = unionedSet.Insert(elem)
	}
	return unionedSet
}

// Intersect 求s和other的交集
func (s *unsafeSet) Intersect(other Set) Set {
	// 将接口类型转换为*unsafeSet类型
	o := other.(*unsafeSet)

	intersection, _ := NewUnsafeSet(-1)
	// loop over smaller s
	if s.Size() < other.Size() {
		for elem := range s.m {
			if other.Contains(elem) {
				_ = intersection.Insert(elem)
			}
		}
	} else {
		for elem := range o.m {
			if s.Contains(elem) {
				_ = intersection.Insert(elem)
			}
		}
	}
	return intersection
}

// Difference 求s - other差集
func (s *unsafeSet) Difference(other Set) Set {
	// 将接口类型转换为*unsafeSet类型
	_ = other.(*unsafeSet)

	difference, _ := NewUnsafeSet(-1)
	for elem := range s.m {
		if !other.Contains(elem) {
			_ = difference.Insert(elem)
		}
	}
	return difference
}

// SymmetricDifference 求该集合s和other的对称差集
// 对称差集：只属于其中一个集合，而不属于另一个集合的元素组成的集合。
func (s *unsafeSet) SymmetricDifference(other Set) Set {
	// 将接口类型转换为*unsafeSet类型
	_ = other.(*unsafeSet)

	aDiff := s.Difference(other)
	bDiff := other.Difference(s)

	return aDiff.Union(bDiff)
}

func (s *unsafeSet) Clear() {
	s, _ = NewUnsafeSet(s.maxSize)
}

func (s *unsafeSet) Remove(value interface{}) {
	delete(s.m, value)
}

func (s *unsafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		for elem := range s.m {
			ch <- elem
		}
		close(ch)
	}()

	return ch
}

func (s *unsafeSet) Iterator() *Iterator {
	iterator, ch, stopCh := newIterator()

	// 开启一个go程对返回的iterator进行监听
	go func() {
	L:
		for elem := range s.m {
			select {
			case <-stopCh:
				break L
			case ch <- elem:
			}
		}
		close(ch)
	}()

	return iterator
}

func (s *unsafeSet) Equal(other Set) bool {
	_ = other.(*unsafeSet)

	if s.Size() != other.Size() {
		return false
	}
	for elem := range s.m {
		if !other.Contains(elem) {
			return false
		}
	}
	return true
}

func (s *unsafeSet) Clone() Set {
	clonedSet, _ := NewUnsafeSet(s.MaxSize())
	for elem := range s.m {
		_ = clonedSet.Insert(elem)
	}
	return clonedSet
}

func (s *unsafeSet) String() string {
	items := make([]string, 0, s.Size())

	for elem := range s.m {
		items = append(items, fmt.Sprintf("%v", elem))
	}
	return fmt.Sprintf("Set{%m}", strings.Join(items, ", "))
}

func (pair OrderedPair) String() string {
	return fmt.Sprintf("(%v, %v)", pair.First, pair.Second)
}

// CartesianProduct 求该集合s和other的笛卡尔积
func (s *unsafeSet) CartesianProduct(other Set) Set {
	o := other.(*unsafeSet)
	cartProduct, _ := NewUnsafeSet(-1)

	for i := range s.m {
		for j := range o.m {
			elem := OrderedPair{First: i, Second: j}
			_ = cartProduct.Insert(elem)
		}
	}

	return cartProduct
}

func (s *unsafeSet) Fill() bool {
	f := false

	if s.maxSize != -1 && len(s.m) == s.maxSize {
		f = true
	}

	return f
}

func (s *unsafeSet) Empty() bool {
	return s.Size() == 0
}

func (s *unsafeSet) Size() int {
	return len(s.m)
}

func (s *unsafeSet) MaxSize() int {
	return s.maxSize
}

func (s *unsafeSet) SetMaxSize(maxSize int) error {
	if maxSize != -1 && maxSize < s.Size() {
		return errors.New("New maxSize is less than current size.")
	}

	s.maxSize = maxSize

	return nil
}

func (s *unsafeSet) CatFromSlice(values []interface{}) error {
	l := len(values)
	if s.maxSize != -1 && s.Size()+l > s.maxSize {
		return errors.New("Not enough free space.")
	}

	for _, value := range values {
		err := s.Insert(value)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *unsafeSet) ToSlice() []interface{} {
	keys := make([]interface{}, 0, s.Size())
	for elem := range s.m {
		keys = append(keys, elem)
	}

	return keys
}

// MarshalJSON 将集合中的所有元素以Json数组的形式返回
func (s *unsafeSet) MarshalJSON() ([]byte, error) {
	items := make([]string, 0, s.Size())

	for elem := range s.m {
		b, err := json.Marshal(elem)
		if err != nil {
			return nil, err
		}

		items = append(items, string(b))
	}

	return []byte(fmt.Sprintf("[%m]", strings.Join(items, ","))), nil
}

// UnmarshalJSON 从给定的Json数组中解析出一个集合,数字将被解析为json.Number
func (s *unsafeSet) UnmarshalJSON(b []byte) error {
	var i []interface{}

	d := json.NewDecoder(bytes.NewReader(b))

	// 使用 UseNumber 方法后，json包会将数字转换成一个内置的 Number 类型（而不是 float64），
	// 这个 Number 类型提供了转换为 int64、float64 等多个方法。
	d.UseNumber()
	err := d.Decode(&i)
	if err != nil {
		return err
	}

	for _, v := range i {
		switch t := v.(type) {
		case []interface{}, map[string]interface{}:
			continue
		default:
			_ = s.Insert(t)
		}
	}

	return nil
}
