package Set

import "sync"

type safeSet struct {
	s unsafeSet
	sync.RWMutex
}

func newThreadSafeSet() safeSet {
	return safeSet{s: newThreadUnsafeSet()}
}

func (set *safeSet) Add(i interface{}) bool {
	set.Lock()
	ret := set.s.Add(i)
	set.Unlock()
	return ret
}

func (set *safeSet) Contains(i ...interface{}) bool {
	set.RLock()
	ret := set.s.Contains(i...)
	set.RUnlock()
	return ret
}

func (set *safeSet) IsSubset(other Set) bool {
	o := other.(*safeSet)

	set.RLock()
	o.RLock()

	ret := set.s.IsSubset(&o.s)
	set.RUnlock()
	o.RUnlock()
	return ret
}

func (set *safeSet) IsProperSubset(other Set) bool {
	o := other.(*safeSet)

	set.RLock()
	defer set.RUnlock()
	o.RLock()
	defer o.RUnlock()

	return set.s.IsProperSubset(&o.s)
}

func (set *safeSet) IsSuperset(other Set) bool {
	return other.IsSubset(set)
}

func (set *safeSet) IsProperSuperset(other Set) bool {
	return other.IsProperSubset(set)
}

func (set *safeSet) Union(other Set) Set {
	o := other.(*safeSet)

	set.RLock()
	o.RLock()

	unsafeUnion := set.s.Union(&o.s).(*unsafeSet)
	ret := &safeSet{s: *unsafeUnion}
	set.RUnlock()
	o.RUnlock()
	return ret
}

func (set *safeSet) Intersect(other Set) Set {
	o := other.(*safeSet)

	set.RLock()
	o.RLock()

	unsafeIntersection := set.s.Intersect(&o.s).(*unsafeSet)
	ret := &safeSet{s: *unsafeIntersection}
	set.RUnlock()
	o.RUnlock()
	return ret
}

func (set *safeSet) Difference(other Set) Set {
	o := other.(*safeSet)

	set.RLock()
	o.RLock()

	unsafeDifference := set.s.Difference(&o.s).(*unsafeSet)
	ret := &safeSet{s: *unsafeDifference}
	set.RUnlock()
	o.RUnlock()
	return ret
}

func (set *safeSet) SymmetricDifference(other Set) Set {
	o := other.(*safeSet)

	set.RLock()
	o.RLock()

	unsafeDifference := set.s.SymmetricDifference(&o.s).(*unsafeSet)
	ret := &safeSet{s: *unsafeDifference}
	set.RUnlock()
	o.RUnlock()
	return ret
}

func (set *safeSet) Clear() {
	set.Lock()
	set.s = newThreadUnsafeSet()
	set.Unlock()
}

func (set *safeSet) Remove(i interface{}) {
	set.Lock()
	delete(set.s, i)
	set.Unlock()
}

func (set *safeSet) Cardinality() int {
	set.RLock()
	defer set.RUnlock()
	return len(set.s)
}

func (set *safeSet) Each(cb func(interface{}) bool) {
	set.RLock()
	for elem := range set.s {
		if cb(elem) {
			break
		}
	}
	set.RUnlock()
}

func (set *safeSet) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		set.RLock()

		for elem := range set.s {
			ch <- elem
		}
		close(ch)
		set.RUnlock()
	}()

	return ch
}

func (set *safeSet) Iterator() *iterator {
	iterator, ch, stopCh := newIterator()

	go func() {
		set.RLock()
	L:
		for elem := range set.s {
			select {
			case <-stopCh:
				break L
			case ch <- elem:
			}
		}
		close(ch)
		set.RUnlock()
	}()

	return iterator
}

func (set *safeSet) Equal(other Set) bool {
	o := other.(*safeSet)

	set.RLock()
	o.RLock()

	ret := set.s.Equal(&o.s)
	set.RUnlock()
	o.RUnlock()
	return ret
}

func (set *safeSet) Clone() Set {
	set.RLock()

	unsafeClone := set.s.Clone().(*unsafeSet)
	ret := &safeSet{s: *unsafeClone}
	set.RUnlock()
	return ret
}

func (set *safeSet) String() string {
	set.RLock()
	ret := set.s.String()
	set.RUnlock()
	return ret
}

func (set *safeSet) PowerSet() Set {
	set.RLock()
	unsafePowerSet := set.s.PowerSet().(*unsafeSet)
	set.RUnlock()

	ret := &safeSet{s: newThreadUnsafeSet()}
	for subset := range unsafePowerSet.Iter() {
		unsafeSubset := subset.(*unsafeSet)
		ret.Add(&safeSet{s: *unsafeSubset})
	}
	return ret
}

func (set *safeSet) Pop() interface{} {
	set.Lock()
	defer set.Unlock()
	return set.s.Pop()
}

func (set *safeSet) CartesianProduct(other Set) Set {
	o := other.(*safeSet)

	set.RLock()
	o.RLock()

	unsafeCartProduct := set.s.CartesianProduct(&o.s).(*unsafeSet)
	ret := &safeSet{s: *unsafeCartProduct}
	set.RUnlock()
	o.RUnlock()
	return ret
}

func (set *safeSet) ToSlice() []interface{} {
	keys := make([]interface{}, 0, set.Cardinality())
	set.RLock()
	for elem := range set.s {
		keys = append(keys, elem)
	}
	set.RUnlock()
	return keys
}

func (set *safeSet) MarshalJSON() ([]byte, error) {
	set.RLock()
	b, err := set.s.MarshalJSON()
	set.RUnlock()

	return b, err
}

func (set *safeSet) UnmarshalJSON(p []byte) error {
	set.RLock()
	err := set.s.UnmarshalJSON(p)
	set.RUnlock()

	return err
}
