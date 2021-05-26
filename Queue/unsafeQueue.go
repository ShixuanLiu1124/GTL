package Queue

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type qNode struct {
	value interface{}
	next  *qNode
	prev  *qNode
}

type unsafeQueue struct {
	size    int
	maxSize int
	head    *qNode
	rail    *qNode
}

func NewUnsafeQueue(maxSize int, values ...interface{}) (*unsafeQueue, error) {
	if maxSize != -1 && len(values) > maxSize {
		return nil, errors.New("Length of values is too long.")
	}

	node := &qNode{
		value: nil,
		next:  nil,
		prev:  nil,
	}

	q := &unsafeQueue{
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

func NewUnsafeQueueWithSlice(maxSize int, values []interface{}) (*unsafeQueue, error) {
	if maxSize != -1 && len(values) > maxSize {
		return nil, errors.New("Length of values is too long.")
	}

	node := &qNode{
		value: nil,
		next:  nil,
		prev:  nil,
	}

	q := &unsafeQueue{
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

func (q *unsafeQueue) Push(value interface{}) error {
	if q.Fill() {
		return errors.New("This queue is fill.")
	}

	node := &qNode{
		value: value,
		next:  nil,
		prev:  q.rail,
	}
	q.rail.next = node
	q.rail = node
	q.size++

	return nil
}

func (q *unsafeQueue) Front() (interface{}, error) {
	if q.Empty() {
		return nil, errors.New("This queue is empty.")
	}

	return q.head.next.value, nil
}

func (q *unsafeQueue) Pop() (interface{}, error) {
	if q.Empty() {
		return nil, errors.New("This queue is empty")
	}

	value := q.head.next.value
	q.head.next = q.head.next.next
	q.size--

	return value, nil
}

/*---------------------------------以下为接口实现---------------------------------------*/

// CatFromSlice 从slice中复制元素到Queue后面
func (q *unsafeQueue) CatFromSlice(values []interface{}) error {
	l := len(values)
	if q.maxSize != -1 && q.size+l > q.maxSize {
		return errors.New("Not enough free space.")
	}

	for _, value := range values {
		err := q.Push(value)
		if err != nil {
			return err
		}
	}
	q.size += l

	return nil
}

func (q *unsafeQueue) Fill() bool {
	f := false
	if q.maxSize != -1 {
		f = q.size == q.maxSize
	}

	return f
}

func (q *unsafeQueue) Empty() bool {
	return q.size == 0
}

func (q *unsafeQueue) Size() int {
	return q.size
}

func (q *unsafeQueue) Clear() {
	q.rail = q.head
	q.head.prev = nil
	q.head.next = nil
	q.head.value = nil
	q.size = 0
}

func (q *unsafeQueue) MaxSize() int {
	return q.maxSize
}

func (q *unsafeQueue) SetMaxSize(maxSize int) error {
	if maxSize != -1 && maxSize < q.size {
		return errors.New("New maxSize is less than current size.")
	}

	q.maxSize = maxSize

	return nil
}

func (q *unsafeQueue) String() string {
	var b strings.Builder
	b.WriteString("unsafeQueue{")

	for p := q.head.next; p != nil; p = p.next {
		if p != q.head.next {
			b.WriteString(", ")
		}
		b.WriteString(fmt.Sprintf("%v", p.value))
	}
	b.WriteString("}")

	return b.String()
}

// ToSlice 将队列以切片形式返回
func (q *unsafeQueue) ToSlice() []interface{} {
	ans := make([]interface{}, q.size)

	for p := q.head.next; p != nil; p = p.next {
		ans = append(ans, p.value)
	}

	return ans
}

// MarshalJSON 将Queue中的所有元素以Json数组的形式返回
func (q *unsafeQueue) MarshalJSON() ([]byte, error) {
	items := make([]string, 0, q.Size())

	for p := q.head.next; p != nil; p = p.next {
		b, err := json.Marshal(p.value)
		if err != nil {
			return nil, err
		}

		items = append(items, string(b))
	}

	return []byte(fmt.Sprintf("[%s]", strings.Join(items, ","))), nil
}

// UnmarshalJSON 从给定的Json数组中解析出一个Queue,数字将被解析为json.Number
func (q *unsafeQueue) UnmarshalJSON(b []byte) error {
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

	return nil
}
