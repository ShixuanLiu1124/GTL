package Set

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

// OrderedPair 表示一个二元组，用于求笛卡尔积
type OrderedPair struct {
	First  interface{}
	Second interface{}
}

type unsafeSet map[interface{}]struct{}

func newUnsafeSet() unsafeSet {
	return make(unsafeSet)
}

// Equal 用来判定两个OrderedPair对象是否相等
func (pair *OrderedPair) Equal(other OrderedPair) bool {
	if pair.First == other.First &&
		pair.Second == other.Second {
		return true
	}

	return false
}

// Add 向集合中添加元素
func (set *unsafeSet) Add(i interface{}) bool {
	_, found := (*set)[i]
	if found {
		return false //False if it existed already
	}

	(*set)[i] = struct{}{}
	return true
}

func (set *unsafeSet) Contains(i ...interface{}) bool {
	for _, val := range i {
		if _, ok := (*set)[val]; !ok {
			return false
		}
	}
	return true
}

// IsSubset 判断other是否是s的子集
func (set *unsafeSet) IsSubset(other Set) bool {
	_ = other.(*unsafeSet)
	if set.Size() > other.Size() {
		return false
	}
	for elem := range *set {
		if !other.Contains(elem) {
			return false
		}
	}
	return true
}

// IsProperSubset 判断other是否是s的真子集
func (set *unsafeSet) IsProperSubset(other Set) bool {
	return set.IsSubset(other) && !set.Equal(other)
}

// IsSuperset 判断other是否是s的超集
func (set *unsafeSet) IsSuperset(other Set) bool {
	return other.IsSubset(set)
}

// IsProperSuperset 判断other是否是s的真超集
func (set *unsafeSet) IsProperSuperset(other Set) bool {
	return set.IsSuperset(other) && !set.Equal(other)
}

// Union 求该集合s和other的并集
func (set *unsafeSet) Union(other Set) Set {
	o := other.(*unsafeSet)

	unionedSet := newUnsafeSet()

	for elem := range *set {
		unionedSet.Add(elem)
	}
	for elem := range *o {
		unionedSet.Add(elem)
	}
	return &unionedSet
}

// Intersect 求s和other的交集
func (set *unsafeSet) Intersect(other Set) Set {
	o := other.(*unsafeSet)

	intersection := newUnsafeSet()
	// loop over smaller set
	if set.Size() < other.Size() {
		for elem := range *set {
			if other.Contains(elem) {
				intersection.Add(elem)
			}
		}
	} else {
		for elem := range *o {
			if set.Contains(elem) {
				intersection.Add(elem)
			}
		}
	}
	return &intersection
}

// Difference 求s - other差集
func (set *unsafeSet) Difference(other Set) Set {
	_ = other.(*unsafeSet)

	difference := newUnsafeSet()
	for elem := range *set {
		if !other.Contains(elem) {
			difference.Add(elem)
		}
	}
	return &difference
}

// SymmetricDifference 求该集合s和other的对称差集
// 对称差集：只属于其中一个集合，而不属于另一个集合的元素组成的集合。
func (set *unsafeSet) SymmetricDifference(other Set) Set {
	_ = other.(*unsafeSet)

	aDiff := set.Difference(other)
	bDiff := other.Difference(set)

	return aDiff.Union(bDiff)
}

func (set *unsafeSet) Clear() {
	*set = newUnsafeSet()
}

func (set *unsafeSet) Remove(i interface{}) {
	delete(*set, i)
}

func (set *unsafeSet) Size() int {
	return len(*set)
}

func (set *unsafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		for elem := range *set {
			ch <- elem
		}
		close(ch)
	}()

	return ch
}

func (set *unsafeSet) Iterator() *Iterator {
	iterator, ch, stopCh := newIterator()

	// 开启一个go程对返回的iterator进行监听
	go func() {
	L:
		for elem := range *set {
			select {
			case <-stopCh:
				break L
			case ch <- elem:
			}
		}
		close(ch)
	}()

	return iterator
}

func (set *unsafeSet) Equal(other Set) bool {
	_ = other.(*unsafeSet)

	if set.Size() != other.Size() {
		return false
	}
	for elem := range *set {
		if !other.Contains(elem) {
			return false
		}
	}
	return true
}

func (set *unsafeSet) Clone() Set {
	clonedSet := newUnsafeSet()
	for elem := range *set {
		clonedSet.Add(elem)
	}
	return &clonedSet
}

func (set *unsafeSet) String() string {
	items := make([]string, 0, len(*set))

	for elem := range *set {
		items = append(items, fmt.Sprintf("%v", elem))
	}
	return fmt.Sprintf("Set{%s}", strings.Join(items, ", "))
}

// String outputs a 2-tuple in the form "(A, B)".
func (pair OrderedPair) String() string {
	return fmt.Sprintf("(%v, %v)", pair.First, pair.Second)
}

// CartesianProduct 求该集合s和other的笛卡尔积
func (set *unsafeSet) CartesianProduct(other Set) Set {
	o := other.(*unsafeSet)
	cartProduct := NewThreadUnsafeSet()

	for i := range *set {
		for j := range *o {
			elem := OrderedPair{First: i, Second: j}
			cartProduct.Add(elem)
		}
	}

	return cartProduct
}

func (set *unsafeSet) ToSlice() []interface{} {
	keys := make([]interface{}, 0, set.Size())
	for elem := range *set {
		keys = append(keys, elem)
	}

	return keys
}

// MarshalJSON 将集合中的所有元素以Json数组的形式返回
func (set *unsafeSet) MarshalJSON() ([]byte, error) {
	items := make([]string, 0, set.Size())

	for elem := range *set {
		b, err := json.Marshal(elem)
		if err != nil {
			return nil, err
		}

		items = append(items, string(b))
	}

	return []byte(fmt.Sprintf("[%s]", strings.Join(items, ","))), nil
}

// UnmarshalJSON 从给定的Json数组中解析出一个集合,数字将被解析为json.Number
func (set *unsafeSet) UnmarshalJSON(b []byte) error {
	var i []interface{}

	d := json.NewDecoder(bytes.NewReader(b))
	d.UseNumber()
	err := d.Decode(&i)
	if err != nil {
		return err
	}

	for _, v := range i {
		switch t := v.(type) {
		case []interface{}, map[string]interface{}:
			continue
		default:
			set.Add(t)
		}
	}

	return nil
}
