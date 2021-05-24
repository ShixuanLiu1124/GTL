package Deque

import (
	"errors"
)

type DQNode struct {
	data interface{}
	next *DQNode
	prev *DQNode
}

type UnsafeDeque struct {
	size    int
	maxSize int
	head    *DQNode
	rail    *DQNode
}

func (q *UnsafeDeque) PushBack(data interface{}) error {
	if q.Fill() {
		return errors.New("This queue is fill.")
	}

	node := &DQNode{
		data: data,
		next: nil,
		prev: q.rail,
	}
	q.rail.next = node
	q.rail = node
	q.size++

	return nil
}

func (q *UnsafeDeque) Front() (interface{}, error) {
	if q.Empty() {
		return nil, errors.New("This queue is empty.")
	}

	return q.head.next.data, nil
}

func (q *UnsafeDeque) PopBack() (interface{}, error) {
	if q.Empty() {
		return nil, errors.New("This queue is empty")
	}

	data := q.head.next.data
	q.head.next = q.head.next.next
	q.size--

	return data, nil
}

/*---------------------------------以下为接口实现---------------------------------------*/

func (q *UnsafeDeque) CopyFromArray(datas []interface{}) error {
	l := len(datas)
	if q.maxSize != -1 && q.size+l > q.maxSize {
		return errors.New("Not enough free space.")
	}

	for a := range datas {
		err := q.PushBack(a)
		if err != nil {
			return err
		}
	}
	q.size += l

	return nil
}

func (q *UnsafeDeque) Fill() bool {
	f := false
	if q.MaxSize() != -1 {
		f = q.Size() == q.MaxSize()
	}

	return f
}

func (q *UnsafeDeque) Empty() bool {
	return q.Size() == 0
}

func (q *UnsafeDeque) Size() int {
	return q.size
}

func (q *UnsafeDeque) MaxSize() int {
	return q.maxSize
}

func (q *UnsafeDeque) SetMaxSize(maxSize int) error {
	if maxSize != -1 && maxSize < q.size {
		return errors.New("New maxSize is less than current size.")
	}

	q.maxSize = maxSize

	return nil
}

func (q *UnsafeDeque) Clear() bool {
	q.rail = q.head
	q.head.prev = nil
	q.head.next = nil
	q.head.data = nil
	q.size = 0

	return true
}
