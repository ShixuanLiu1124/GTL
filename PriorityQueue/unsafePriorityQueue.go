package PriorityQueue

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type unsafePriorityQueue struct {
	maxSize int
	s       []interface{}
	less    func(interface{}, interface{}) bool
}

func (q *unsafePriorityQueue) Push(value interface{}) error {
	if q.Fill() {
		return errors.New("This queue is fill.")
	}

	q.s = append(q.s, value)
	q.Fix(0)

	return nil
}

func (q *unsafePriorityQueue) Pop() (interface{}, error) {
	if q.Empty() {
		return nil, errors.New("This queue is empty")
	}

	value := q.s[len(q.s)-1]
	q.s = q.s[:len(q.s)-1]

	q.Fix(0)

	return value, nil
}

func (q *unsafePriorityQueue) Top() (interface{}, error) {
	if q.Empty() {
		return nil, errors.New("This priority queue is empty")
	}

	return q.s[0], nil
}

func (q *unsafePriorityQueue) Fix(index int) {
	// TODO 完成对堆的调整
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

func (q *unsafePriorityQueue) SetFunc(less func(interface{}, interface{}) bool) {
	q.less = less
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

	q.Fix(0)

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

	q.Fix(0)

	return nil
}
