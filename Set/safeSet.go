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
	"sync"
)

type safeSet struct {
	us *unsafeSet
	sync.RWMutex
}

func NewSafeSet(maxSize int, values ...interface{}) (*safeSet, error) {
	s, err := NewUnsafeSet(maxSize, values...)
	if err != nil {
		return nil, err
	}

	return &safeSet{
		us:      s,
		RWMutex: sync.RWMutex{},
	}, nil
}

func NewSafeSetWithSlice(maxSize int, values []interface{}) (*safeSet, error) {
	s, err := NewUnsafeSetWithSlice(maxSize, values)
	if err != nil {
		return nil, err
	}

	return &safeSet{
		us:      s,
		RWMutex: sync.RWMutex{},
	}, nil
}

func (set *safeSet) Insert(value interface{}) error {
	set.Lock()
	err := set.us.Insert(value)
	set.Unlock()
	return err
}

func (set *safeSet) Contains(values ...interface{}) bool {
	set.RLock()
	ret := set.us.Contains(values...)
	set.RUnlock()
	return ret
}

func (set *safeSet) IsSubset(other Set) bool {
	o := other.(*safeSet)

	set.RLock()
	o.RLock()

	ret := set.us.IsSubset(o.us)
	set.RUnlock()
	o.RUnlock()

	return ret
}

func (set *safeSet) IsProperSubset(other Set) bool {
	o := other.(*safeSet)

	set.RLock()
	defer set.RUnlock()
	o.RLock()
	defer o.RUnlock()

	return set.us.IsProperSubset(o.us)
}

func (set *safeSet) IsSuperset(other Set) bool {
	return other.IsSubset(set)
}

func (set *safeSet) IsProperSuperset(other Set) bool {
	return other.IsProperSubset(set)
}

func (set *safeSet) Union(other Set) Set {
	o := other.(*safeSet)

	set.RLock()
	o.RLock()

	unsafeUnion := set.us.Union(o.us).(*unsafeSet)
	ret := &safeSet{
		us:      unsafeUnion,
		RWMutex: sync.RWMutex{},
	}
	set.RUnlock()
	o.RUnlock()

	return ret
}

func (set *safeSet) Intersect(other Set) Set {
	o := other.(*safeSet)

	set.RLock()
	o.RLock()

	unsafeIntersection := set.us.Intersect(o.us).(*unsafeSet)
	ret := &safeSet{
		us:      unsafeIntersection,
		RWMutex: sync.RWMutex{},
	}
	set.RUnlock()
	o.RUnlock()

	return ret
}

func (set *safeSet) Difference(other Set) Set {
	o := other.(*safeSet)

	set.RLock()
	o.RLock()

	unsafeDifference := set.us.Difference(o.us).(*unsafeSet)
	ret := &safeSet{
		us:      unsafeDifference,
		RWMutex: sync.RWMutex{},
	}
	set.RUnlock()
	o.RUnlock()
	return ret
}

func (set *safeSet) SymmetricDifference(other Set) Set {
	o := other.(*safeSet)

	set.RLock()
	o.RLock()

	unsafeDifference := set.us.SymmetricDifference(o.us).(*unsafeSet)
	ret := &safeSet{
		us:      unsafeDifference,
		RWMutex: sync.RWMutex{},
	}
	set.RUnlock()
	o.RUnlock()

	return ret
}

func (set *safeSet) Remove(value interface{}) {
	set.Lock()
	delete(set.us.m, value)
	set.Unlock()
}

func (set *safeSet) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		set.RLock()

		for elem := range set.us.m {
			ch <- elem
		}
		close(ch)
		set.RUnlock()
	}()

	return ch
}

func (set *safeSet) Iterator() *Iterator {
	iterator, ch, stopCh := newIterator()

	go func() {
		set.RLock()
	L:
		for elem := range set.us.m {
			select {
			case <-stopCh:
				break L
			case ch <- elem:
			}
		}
		close(ch)
		set.RUnlock()
	}()

	return iterator
}

func (set *safeSet) Equal(other Set) bool {
	o := other.(*safeSet)

	set.RLock()
	o.RLock()

	ret := set.us.Equal(o.us)
	set.RUnlock()
	o.RUnlock()

	return ret
}

func (set *safeSet) Clone() Set {
	set.RLock()
	defer set.RUnlock()

	unsafeClone := set.us.Clone().(*unsafeSet)
	return &safeSet{
		us:      unsafeClone,
		RWMutex: sync.RWMutex{},
	}
}

func (set *safeSet) String() string {
	set.RLock()
	defer set.RUnlock()

	return set.us.String()
}

func (set *safeSet) CartesianProduct(other Set) Set {
	o := other.(*safeSet)

	set.RLock()
	o.RLock()

	unsafeCartProduct := set.us.CartesianProduct(o.us).(*unsafeSet)
	ret := &safeSet{
		us:      unsafeCartProduct,
		RWMutex: sync.RWMutex{},
	}
	set.RUnlock()
	o.RUnlock()

	return ret
}

func (set *safeSet) Clear() {
	set.Lock()
	set.us, _ = NewUnsafeSet(set.MaxSize())
	set.Unlock()
}

func (set *safeSet) Fill() bool {
	set.RLock()
	defer set.RUnlock()

	return set.Fill()
}

func (set *safeSet) Empty() bool {
	set.RLock()
	defer set.RUnlock()

	return set.us.Empty()
}

func (set *safeSet) Size() int {
	set.RLock()
	defer set.RUnlock()

	return set.us.Size()
}

func (set *safeSet) MaxSize() int {
	set.RLock()
	defer set.RUnlock()

	return set.us.MaxSize()
}

func (set *safeSet) SetMaxSize(maxSize int) error {
	set.Lock()
	defer set.Unlock()

	return set.us.SetMaxSize(maxSize)
}

func (set *safeSet) CatFromSlice(values []interface{}) error {
	set.Lock()
	defer set.Unlock()

	return set.us.CatFromSlice(values)
}

func (set *safeSet) ToSlice() []interface{} {
	keys := make([]interface{}, 0, set.Size())
	set.RLock()
	for elem := range set.us.m {
		keys = append(keys, elem)
	}
	set.RUnlock()
	return keys
}

func (set *safeSet) MarshalJSON() ([]byte, error) {
	set.RLock()
	b, err := set.us.MarshalJSON()
	set.RUnlock()

	return b, err
}

func (set *safeSet) UnmarshalJSON(p []byte) error {
	set.Lock()
	err := set.us.UnmarshalJSON(p)
	set.Unlock()

	return err
}
