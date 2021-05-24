package Queue

import "errors"

type QNode struct {
	value interface{}
	next  *QNode
	prev  *QNode
}

type UnsafeQueue struct {
	size    int
	maxSize int
	head    *QNode
	rail    *QNode
}

func New(maxSize int, values ...interface{}) (*UnsafeQueue, error) {
	if maxSize != -1 && len(values) > maxSize {
		return nil, errors.New("Length of values is too long.")
	}

	node := &QNode{
		value: nil,
		next:  nil,
		prev:  nil,
	}

	q := &UnsafeQueue{
		size:    0,
		maxSize: maxSize,
		head:    node,
		rail:    node,
	}

	for _, value := range values {
		err := q.Push(value)
		if err != nil {
			return nil, err
		}
	}

	return q, nil
}

func (q *UnsafeQueue) Push(value interface{}) error {
	if q.Fill() {
		return errors.New("This queue is fill.")
	}

	node := &QNode{
		value: value,
		next:  nil,
		prev:  q.rail,
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

	return q.head.next.value, nil
}

func (q *UnsafeQueue) Pop() (interface{}, error) {
	if q.Empty() {
		return nil, errors.New("This queue is empty")
	}

	value := q.head.next.value
	q.head.next = q.head.next.next
	q.size--

	return value, nil
}

/*---------------------------------以下为接口实现---------------------------------------*/

func (q *UnsafeQueue) CopyFromArray(values []interface{}) error {
	l := len(values)
	if q.maxSize != -1 && q.size+l > q.maxSize {
		return errors.New("Not enough free space.")
	}

	for a := range values {
		err := q.Push(a)
		if err != nil {
			return err
		}
	}
	q.size += l

	return nil
}

func (q *UnsafeQueue) Fill() bool {
	f := false
	if q.maxSize != -1 {
		f = q.size == q.maxSize
	}

	return f
}

func (q *UnsafeQueue) Empty() bool {
	return q.size == 0
}

func (q *UnsafeQueue) Size() int {
	return q.size
}

func (q *UnsafeQueue) Clear() bool {
	q.rail = q.head
	q.head.prev = nil
	q.head.next = nil
	q.head.value = nil
	q.size = 0

	return true
}

func (q *UnsafeQueue) MaxSize() int {
	return q.maxSize
}

func (q *UnsafeQueue) SetMaxSize(maxSize int) error {
	if maxSize != -1 && maxSize < q.size {
		return errors.New("New maxSize is less than current size.")
	}

	q.maxSize = maxSize

	return nil
}
