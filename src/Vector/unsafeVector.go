package Vector

import (
	"errors"
	"fmt"
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
		s:       values,
		maxSize: maxSize,
	}

	return v, nil
}

func (v *unsafeVector) PushBack(value interface{}) error {
	if v.Fill() {
		return errors.New("This queue is fill.")
	}

	v.s = append(v.s, value)

	return nil
}

// PopBack 弹出最后一个元素
func (v *unsafeVector) PopBack() (interface{}, error) {
	if v.Empty() {
		return nil, errors.New("This Vector is empty.")
	}

	value := v.s[len(v.s)-1]
	v.s = v.s[:len(v.s)-1]

	return value, nil
}

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
	// TODO: 设定阈值
	// 小于阈值时，使用nil赋值
	v.s = nil

	// 大于阈值时，使用引用赋值
	v.s = v.s[:0]
}

func (v *unsafeVector) String() string {
	return fmt.Sprintf("%v", v.s)
}

func (v *unsafeVector) CopyFromArray(values []interface{}) error {
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
	return v.s
}
