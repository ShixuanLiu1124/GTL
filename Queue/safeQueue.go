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

package Queue

import (
	"sync"
)

type safeQueue struct {
	uq *unsafeQueue
	m  *sync.RWMutex
}

func NewSafeQueue(maxSize int, values ...interface{}) (*safeQueue, error) {
	q, err := NewUnsafeQueue(maxSize, values)

	return &safeQueue{
		uq: q,
		m:  new(sync.RWMutex),
	}, err
}

func NewSafeQueueWithSlice(maxSize int, values []interface{}) (*safeQueue, error) {
	q, err := NewUnsafeQueueWithSlice(maxSize, values)

	return &safeQueue{
		uq: q,
		m:  new(sync.RWMutex),
	}, err
}

func (q *safeQueue) Push(value interface{}) error {
	q.m.Lock()
	defer q.m.Unlock()

	return q.uq.Push(value)
}

func (q *safeQueue) Front() (interface{}, error) {
	q.m.RLock()
	defer q.m.RUnlock()

	return q.uq.Front()
}

func (q *safeQueue) Pop() (interface{}, error) {
	q.m.Lock()
	defer q.m.Unlock()

	return q.uq.Pop()
}

func (q *safeQueue) Fill() bool {
	q.m.RLock()
	defer q.m.RUnlock()

	return q.uq.Fill()
}

func (q *safeQueue) Empty() bool {
	q.m.RLock()
	defer q.m.RUnlock()

	return q.uq.Empty()
}

func (q *safeQueue) Size() int {
	q.m.RLock()
	defer q.m.RUnlock()

	return q.uq.Size()
}

func (q *safeQueue) MaxSize() int {
	q.m.RLock()
	defer q.m.RUnlock()

	return q.uq.MaxSize()
}

func (q *safeQueue) SetMaxSize(i int) error {
	q.m.Lock()
	defer q.m.Unlock()

	return q.uq.SetMaxSize(i)
}

func (q *safeQueue) Clear() {
	q.m.Lock()
	defer q.m.Unlock()

	q.uq.Clear()
}

func (q *safeQueue) String() string {
	q.m.RLock()
	defer q.m.RUnlock()

	return q.uq.String()
}

func (q *safeQueue) CatFromSlice(values []interface{}) error {
	q.m.Lock()
	defer q.m.Unlock()

	return q.uq.CatFromSlice(values)
}

func (q *safeQueue) ToSlice() []interface{} {
	q.m.RLock()
	defer q.m.RUnlock()

	return q.uq.ToSlice()
}

func (q *safeQueue) MarshalJSON() ([]byte, error) {
	q.m.RLock()
	defer q.m.RUnlock()

	return q.uq.MarshalJSON()
}

func (q *safeQueue) UnmarshalJSON(b []byte) error {
	q.m.Lock()
	defer q.m.Unlock()

	return q.uq.UnmarshalJSON(b)
}
