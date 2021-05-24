package Set

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// An OrderedPair represents a 2-tuple of values.
type OrderedPair struct {
	First  interface{}
	Second interface{}
}

type UnsafeSet map[interface{}]struct{}

func newThreadUnsafeSet() UnsafeSet {
	return make(UnsafeSet)
}

// Equal says whether two 2-tuples contain the same values in the same order.
func (pair *OrderedPair) Equal(other OrderedPair) bool {
	if pair.First == other.First &&
		pair.Second == other.Second {
		return true
	}

	return false
}

func (set *UnsafeSet) Add(i interface{}) bool {
	_, found := (*set)[i]
	if found {
		return false //False if it existed already
	}

	(*set)[i] = struct{}{}
	return true
}

func (set *UnsafeSet) Contains(i ...interface{}) bool {
	for _, val := range i {
		if _, ok := (*set)[val]; !ok {
			return false
		}
	}
	return true
}

func (set *UnsafeSet) IsSubset(other Set) bool {
	_ = other.(*UnsafeSet)
	if set.Cardinality() > other.Cardinality() {
		return false
	}
	for elem := range *set {
		if !other.Contains(elem) {
			return false
		}
	}
	return true
}

func (set *UnsafeSet) IsProperSubset(other Set) bool {
	return set.IsSubset(other) && !set.Equal(other)
}

func (set *UnsafeSet) IsSuperset(other Set) bool {
	return other.IsSubset(set)
}

func (set *UnsafeSet) IsProperSuperset(other Set) bool {
	return set.IsSuperset(other) && !set.Equal(other)
}

func (set *UnsafeSet) Union(other Set) Set {
	o := other.(*UnsafeSet)

	unionedSet := newThreadUnsafeSet()

	for elem := range *set {
		unionedSet.Add(elem)
	}
	for elem := range *o {
		unionedSet.Add(elem)
	}
	return &unionedSet
}

func (set *UnsafeSet) Intersect(other Set) Set {
	o := other.(*UnsafeSet)

	intersection := newThreadUnsafeSet()
	// loop over smaller set
	if set.Cardinality() < other.Cardinality() {
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

func (set *UnsafeSet) Difference(other Set) Set {
	_ = other.(*UnsafeSet)

	difference := newThreadUnsafeSet()
	for elem := range *set {
		if !other.Contains(elem) {
			difference.Add(elem)
		}
	}
	return &difference
}

func (set *UnsafeSet) SymmetricDifference(other Set) Set {
	_ = other.(*UnsafeSet)

	aDiff := set.Difference(other)
	bDiff := other.Difference(set)
	return aDiff.Union(bDiff)
}

func (set *UnsafeSet) Clear() {
	*set = newThreadUnsafeSet()
}

func (set *UnsafeSet) Remove(i interface{}) {
	delete(*set, i)
}

func (set *UnsafeSet) Cardinality() int {
	return len(*set)
}

func (set *UnsafeSet) Each(cb func(interface{}) bool) {
	for elem := range *set {
		if cb(elem) {
			break
		}
	}
}

func (set *UnsafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		for elem := range *set {
			ch <- elem
		}
		close(ch)
	}()

	return ch
}

func (set *UnsafeSet) Iterator() *Iterator {
	iterator, ch, stopCh := newIterator()

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

func (set *UnsafeSet) Equal(other Set) bool {
	_ = other.(*UnsafeSet)

	if set.Cardinality() != other.Cardinality() {
		return false
	}
	for elem := range *set {
		if !other.Contains(elem) {
			return false
		}
	}
	return true
}

func (set *UnsafeSet) Clone() Set {
	clonedSet := newThreadUnsafeSet()
	for elem := range *set {
		clonedSet.Add(elem)
	}
	return &clonedSet
}

func (set *UnsafeSet) String() string {
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

func (set *UnsafeSet) Pop() interface{} {
	for item := range *set {
		delete(*set, item)
		return item
	}
	return nil
}

func (set *UnsafeSet) PowerSet() Set {
	powSet := NewThreadUnsafeSet()
	nullset := newThreadUnsafeSet()
	powSet.Add(&nullset)

	for es := range *set {
		u := newThreadUnsafeSet()
		j := powSet.Iter()
		for er := range j {
			p := newThreadUnsafeSet()
			if reflect.TypeOf(er).Name() == "" {
				k := er.(*UnsafeSet)
				for ek := range *(k) {
					p.Add(ek)
				}
			} else {
				p.Add(er)
			}
			p.Add(es)
			u.Add(&p)
		}

		powSet = powSet.Union(&u)
	}

	return powSet
}

func (set *UnsafeSet) CartesianProduct(other Set) Set {
	o := other.(*UnsafeSet)
	cartProduct := NewThreadUnsafeSet()

	for i := range *set {
		for j := range *o {
			elem := OrderedPair{First: i, Second: j}
			cartProduct.Add(elem)
		}
	}

	return cartProduct
}

func (set *UnsafeSet) ToSlice() []interface{} {
	keys := make([]interface{}, 0, set.Cardinality())
	for elem := range *set {
		keys = append(keys, elem)
	}

	return keys
}

// MarshalJSON creates a JSON array from the set, it marshals all elements
func (set *UnsafeSet) MarshalJSON() ([]byte, error) {
	items := make([]string, 0, set.Cardinality())

	for elem := range *set {
		b, err := json.Marshal(elem)
		if err != nil {
			return nil, err
		}

		items = append(items, string(b))
	}

	return []byte(fmt.Sprintf("[%s]", strings.Join(items, ","))), nil
}

// UnmarshalJSON recreates a set from a JSON array, it only decodes
// primitive types. Numbers are decoded as json.Number.
func (set *UnsafeSet) UnmarshalJSON(b []byte) error {
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
