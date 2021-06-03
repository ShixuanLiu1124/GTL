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

import "sync"

type safeDeque struct {
	uq *unsafeDeque
	m  *sync.RWMutex
}

func NewSafeDeque(maxSize int, values ...interface{}) (*safeDeque, error) {
	q, err := NewUnsafeDeque(maxSize, values)

	return &safeDeque{
		uq: q,
		m:  new(sync.RWMutex),
	}, err
}

func NewSafeDequeWithSlice(maxSize int, values []interface{}) (*safeDeque, error) {
	q, err := NewUnsafeDequeWithSlice(maxSize, values)

	return &safeDeque{
		uq: q,
		m:  new(sync.RWMutex),
	}, err
}

func (q *safeDeque) PushFront(value interface{}) error {
	q.m.Lock()
	defer q.m.Unlock()

	return q.uq.PushFront(value)
}

func (q *safeDeque) PushBack(value interface{}) error {
	q.m.Lock()
	defer q.m.Unlock()

	return q.uq.PushBack(value)
}

func (q *safeDeque) Front() (interface{}, error) {
	q.m.RLock()
	defer q.m.RUnlock()

	return q.uq.Front()
}

func (q *safeDeque) Back() (interface{}, error) {
	q.m.RLock()
	defer q.m.RUnlock()

	return q.uq.Back()
}

func (q *safeDeque) PopFront() (interface{}, error) {
	q.m.Lock()
	defer q.m.Unlock()

	return q.uq.PopFront()
}

func (q *safeDeque) PopBack() (interface{}, error) {
	q.m.Lock()
	defer q.m.Unlock()

	return q.uq.PopBack()
}

func (q *safeDeque) Fill() bool {
	q.m.RLock()
	defer q.m.RUnlock()

	return q.uq.Fill()
}

func (q *safeDeque) Empty() bool {
	q.m.RLock()
	defer q.m.RUnlock()

	return q.uq.Empty()
}

func (q *safeDeque) Size() int {
	q.m.RLock()
	defer q.m.RUnlock()

	return q.uq.Size()
}

func (q *safeDeque) MaxSize() int {
	q.m.RLock()
	defer q.m.RUnlock()

	return q.uq.MaxSize()
}

func (q *safeDeque) SetMaxSize(i int) error {
	q.m.Lock()
	defer q.m.Unlock()

	return q.uq.SetMaxSize(i)
}

func (q *safeDeque) Clear() {
	q.m.Lock()
	defer q.m.Unlock()

	q.uq.Clear()
}

func (q *safeDeque) String() string {
	q.m.RLock()
	defer q.m.RUnlock()

	return q.uq.String()
}

func (q *safeDeque) CatFromSlice(values []interface{}) error {
	q.m.Lock()
	defer q.m.Unlock()

	return q.uq.CatFromSlice(values)
}

func (q *safeDeque) ToSlice() []interface{} {
	q.m.RLock()
	defer q.m.RUnlock()

	return q.uq.ToSlice()
}

func (q *safeDeque) MarshalJSON() ([]byte, error) {
	q.m.RLock()
	defer q.m.RUnlock()

	return q.uq.MarshalJSON()
}

func (q *safeDeque) UnmarshalJSON(b []byte) error {
	q.m.Lock()
	defer q.m.Unlock()

	return q.uq.UnmarshalJSON(b)
}
