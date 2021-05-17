package Queue

import "errors"

type QNode struct {
	data interface{}
	next *QNode
	prev *QNode
}

type Queue struct {
	size    int
	head    *QNode
	rail    *QNode
	maxSize int
}

func (q *Queue) Init(maxSize int) *Queue {
	q = new(Queue)
	node := &QNode{
		data: nil,
		next: nil,
		prev: nil,
	}
	q.head = node
	q.rail = node
	q.maxSize = maxSize
	q.size = 0

	return q
}

func (q *Queue) Fill() bool {
	f := false
	if q.MaxSize() != -1 {
		f = q.Size() == q.MaxSize()
	}

	return f
}

func (q *Queue) Empty() bool {
	return q.Size() == 0
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) ToString() string {
	// TODO: ToString method
	return ""
}

func (q *Queue) Clear() bool {
	q.rail = q.head
	q.head.prev = nil
	q.head.next = nil
	q.head.data = nil
	q.size = 0

	return true
}

func (q *Queue) MaxSize() int {
	return q.maxSize
}

func (q *Queue) Push(data interface{}) error {
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

func (q *Queue) Front() (interface{}, error) {
	if q.Empty() {
		return nil, errors.New("This queue is empty.")
	}

	return q.head.next.data, nil
}

func (q *Queue) Pop() (interface{}, error) {
	if q.Empty() {
		return nil, errors.New("This queue is empty")
	}

	data := q.head.next.data
	q.head.next = q.head.next.next
	q.size--

	return data, nil
}
