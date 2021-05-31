package Vector

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type unsafeVector struct {
	s       []interface{}
	maxSize int
}

func NewUnsafeVector(maxSize int, values ...interface{}) (*unsafeVector, error) {
	if maxSize != -1 && len(values) > maxSize {
		return nil, errors.New("Length of values is too long.")
	}

	v := &unsafeVector{
		s:       values,
		maxSize: maxSize,
	}

	return v, nil
}

func NewUnsafeVectorWithSlice(maxSize int, values []interface{}) (*unsafeVector, error) {
	if maxSize != -1 && len(values) > maxSize {
		return nil, errors.New("Length of values is too long.")
	}

	v := &unsafeVector{
		s:       []interface{}{},
		maxSize: maxSize,
	}

	for _, value := range values {
		v.s = append(v.s, value)
	}

	return v, nil
}

// PushBack 从vector后方加入元素
func (v *unsafeVector) PushBack(value interface{}) error {
	if v.Fill() {
		return errors.New("This vector is fill.")
	}

	v.s = append(v.s, value)

	return nil
}

// PopBack 弹出最后一个元素
func (v *unsafeVector) PopBack() (interface{}, error) {
	if v.Empty() {
		return nil, errors.New("This vector is empty.")
	}

	value := v.s[len(v.s)-1]
	v.s = v.s[:len(v.s)-1]

	return value, nil
}

func (v *unsafeVector) Set(index, value int) error {
	b, err := v.At(index)
	if err != nil {
		return err
	}
	if b != value {
		v.s[index] = value
	}
	return nil
}

// At 返回位于index处的元素
func (v *unsafeVector) At(index int) (interface{}, error) {
	if v.Size() < index+1 {
		return nil, errors.New("Index out of bounds")
	}

	return v.s[index], nil
}

// Remove 删除下标位于区间[start, end)之间的元素
func (v *unsafeVector) Remove(start, end int) error {
	if start < 0 || end > v.Size() {
		return errors.New("Index out of bounds.")
	}

	v.s = append(v.s[:start], v.s[end:]...)

	return nil
}

// Find 使用二分查找技术查找元素下标，less是比较函数，用于比较value1是否小于value2
func (v *unsafeVector) Find(
	value interface{},
	less func(interface{}, interface{}) bool,
) int {
	start := 0
	end := v.Size()
	pos := -1

	for start < end {
		mid := start + (end-start)/2
		temp, err := v.At(mid)
		if err != nil {
			fmt.Println(err)
			return -1
		}

		if temp == value {
			pos = mid
			break
		} else if less(temp, value) {
			start = mid + 1
		} else {
			end = mid
		}
	}

	return pos
}

/*---------------------------------以下为接口实现---------------------------------------*/

func (v *unsafeVector) Fill() bool {
	f := false

	if v.maxSize != -1 && len(v.s) == v.maxSize {
		f = true
	}

	return f
}

func (v *unsafeVector) Empty() bool {
	return len(v.s) == 0
}

func (v *unsafeVector) Size() int {
	return len(v.s)
}

func (v *unsafeVector) MaxSize() int {
	return v.maxSize
}

func (v *unsafeVector) SetMaxSize(maxSize int) error {
	if maxSize != -1 && maxSize < v.Size() {
		return errors.New("New maxSize is less than current size.")
	}

	v.maxSize = maxSize

	return nil
}

func (v *unsafeVector) Clear() {
	v.s = nil
}

func (v *unsafeVector) String() string {
	return fmt.Sprintf("%v", v.s)
}

// CatFromSlice 从slice中复制元素到Vector后面
func (v *unsafeVector) CatFromSlice(values []interface{}) error {
	l := len(values)
	if v.maxSize != -1 && v.Size()+l > v.maxSize {
		return errors.New("Not enough free space.")
	}

	for _, value := range values {
		err := v.PushBack(value)
		if err != nil {
			return err
		}
	}

	return nil
}

func (v *unsafeVector) ToSlice() []interface{} {
	// 切片直接指向存储空间，所以要复制到临时变量中再返回
	b := make([]interface{}, len(v.s))
	copy(b, v.s)

	return b
}

// MarshalJSON 将Vector中的所有元素以Json数组的形式返回
func (v *unsafeVector) MarshalJSON() ([]byte, error) {
	items := make([]string, 0, v.Size())

	for elem := range v.s {
		b, err := json.Marshal(elem)
		if err != nil {
			return nil, err
		}

		items = append(items, string(b))
	}

	return []byte(fmt.Sprintf("[%s]", strings.Join(items, ","))), nil
}

// UnmarshalJSON 从给定的Json数组中解析出一个Vector,数字将被解析为json.Number
func (v *unsafeVector) UnmarshalJSON(b []byte) error {
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
			err = v.PushBack(t)
			if err != nil {
				fmt.Println(err)
				return err
			}
		}
	}

	return nil
}
