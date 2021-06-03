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

import "sync"

type safePriorityQueue struct {
	uq *unsafePriorityQueue
	m  sync.RWMutex
}

func NewSafePriorityQueue(maxSize int, values ...interface{}) (*safePriorityQueue, error) {
	q, err := NewUnsafePriorityQueue(maxSize, values)

	return &safePriorityQueue{
		uq: q,
		m:  sync.RWMutex{},
	}, err
}

func NewSafePriorityQueueWithSlice(maxSize int, values []interface{}) (*safePriorityQueue, error) {
	q, err := NewUnsafePriorityQueueWithSlice(maxSize, values)

	return &safePriorityQueue{
		uq: q,
		m:  sync.RWMutex{},
	}, err
}

func (q *safePriorityQueue) Push(value interface{}) error {
	q.m.Lock()
	defer q.m.Unlock()

	return q.uq.Push(value)
}

func (q *safePriorityQueue) Pop() (interface{}, error) {
	q.m.Lock()
	defer q.m.Unlock()

	return q.uq.Pop()
}

func (q *safePriorityQueue) Top() (interface{}, error) {
	q.m.RLock()
	defer q.m.RUnlock()

	return q.Top()
}

func (q *safePriorityQueue) fix(index int) {
	q.m.Lock()
	defer q.m.Unlock()

	q.uq.fix(index)
}

func (q *safePriorityQueue) SetFunc(less func(interface{}, interface{}) bool) {
	q.m.Lock()
	defer q.m.Unlock()

	q.uq.SetFunc(less)
}

func (q *safePriorityQueue) Fill() bool {
	q.m.RLock()
	defer q.m.RUnlock()

	return q.uq.Fill()
}

func (q *safePriorityQueue) Empty() bool {
	q.m.RLock()
	defer q.m.RUnlock()

	return q.uq.Empty()
}

func (q *safePriorityQueue) Size() int {
	q.m.RLock()
	defer q.m.RUnlock()

	return q.uq.Size()
}

func (q *safePriorityQueue) MaxSize() int {
	q.m.RLock()
	defer q.m.RUnlock()

	return q.uq.MaxSize()
}

func (q *safePriorityQueue) SetMaxSize(maxSize int) error {
	q.m.Lock()
	defer q.m.Unlock()

	return q.uq.SetMaxSize(maxSize)
}

func (q *safePriorityQueue) Clear() {
	q.m.Lock()
	defer q.m.Unlock()

	q.uq.Clear()
}

func (q *safePriorityQueue) String() string {
	q.m.RLock()
	defer q.m.RUnlock()

	return q.uq.String()
}

func (q *safePriorityQueue) CatFromSlice(values []interface{}) error {
	q.m.Lock()
	defer q.m.Unlock()

	return q.uq.CatFromSlice(values)
}

func (q *safePriorityQueue) ToSlice() []interface{} {
	q.m.Lock()
	defer q.m.Unlock()

	return q.uq.ToSlice()
}

func (q *safePriorityQueue) MarshalJSON() ([]byte, error) {
	q.m.RLock()
	defer q.m.RUnlock()

	return q.uq.MarshalJSON()
}

func (q *safePriorityQueue) UnmarshalJSON(b []byte) error {
	q.m.Lock()
	defer q.m.Unlock()

	return q.uq.UnmarshalJSON(b)
}
