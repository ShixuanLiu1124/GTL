package Queue

import "errors"

type QNode struct {
	data interface{}
	next *QNode
	prev *QNode
}

type UnsafeQueue struct {
	size    int
	maxSize int
	head    *QNode
	rail    *QNode
}

func New(maxSize int) *UnsafeQueue {
	node := &QNode{
		data: nil,
		next: nil,
		prev: nil,
	}

	return &UnsafeQueue{
		size:    0,
		maxSize: maxSize,
		head:    node,
		rail:    node,
	}
}

func (q *UnsafeQueue) CopyFromArray(datas []interface{}) error {
	l := len(datas)
	if q.maxSize != -1 && q.size+l > q.maxSize {
		return errors.New("Not enough free space.")
	}

	for a := range datas {
		err := q.Push(a)
		if err != nil {
			return err
		}
	}
	q.size += l

	return nil
}

func (q *UnsafeQueue) SetMaxSize(maxSize int) error {
	if maxSize != -1 && maxSize < q.size {
		return errors.New("New maxSize is less than current size.")
	}

	q.maxSize = maxSize

	return nil
}

func (q *UnsafeQueue) Push(data interface{}) error {
	if q.Fill() {
		return errors.New("This queue is fill.")
	}

	node := &QNode{
		data: data,
		next: nil,
		prev: q.rail,
	}
	q.rail.next = node
	q.rail = node
	q.size++

	return nil
}

func (q *UnsafeQueue) Front() (interface{}, error) {
	if q.Empty() {
		return nil, errors.New("This queue is empty.")
	}

	return q.head.next.data, nil
}

func (q *UnsafeQueue) Pop() (interface{}, error) {
	if q.Empty() {
		return nil, errors.New("This queue is empty")
	}

	data := q.head.next.data
	q.head.next = q.head.next.next
	q.size--

	return data, nil
}

func (q *UnsafeQueue) Fill() bool {
	f := false
	if q.MaxSize() != -1 {
		f = q.Size() == q.MaxSize()
	}

	return f
}

func (q *UnsafeQueue) Empty() bool {
	return q.Size() == 0
}

func (q *UnsafeQueue) Size() int {
	return q.size
}

func (q *UnsafeQueue) ToString() string {
	// TODO: ToString method
	return ""
}

func (q *UnsafeQueue) Clear() bool {
	q.rail = q.head
	q.head.prev = nil
	q.head.next = nil
	q.head.data = nil
	q.size = 0

	return true
}

func (q *UnsafeQueue) MaxSize() int {
	return q.maxSize
}
