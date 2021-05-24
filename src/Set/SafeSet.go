package Set

import "sync"

type SafeSet struct {
	s UnsafeSet
	sync.RWMutex
}

func newThreadSafeSet() SafeSet {
	return SafeSet{s: newThreadUnsafeSet()}
}

func (set *SafeSet) Add(i interface{}) bool {
	set.Lock()
	ret := set.s.Add(i)
	set.Unlock()
	return ret
}

func (set *SafeSet) Contains(i ...interface{}) bool {
	set.RLock()
	ret := set.s.Contains(i...)
	set.RUnlock()
	return ret
}

func (set *SafeSet) IsSubset(other Set) bool {
	o := other.(*SafeSet)

	set.RLock()
	o.RLock()

	ret := set.s.IsSubset(&o.s)
	set.RUnlock()
	o.RUnlock()
	return ret
}

func (set *SafeSet) IsProperSubset(other Set) bool {
	o := other.(*SafeSet)

	set.RLock()
	defer set.RUnlock()
	o.RLock()
	defer o.RUnlock()

	return set.s.IsProperSubset(&o.s)
}

func (set *SafeSet) IsSuperset(other Set) bool {
	return other.IsSubset(set)
}

func (set *SafeSet) IsProperSuperset(other Set) bool {
	return other.IsProperSubset(set)
}

func (set *SafeSet) Union(other Set) Set {
	o := other.(*SafeSet)

	set.RLock()
	o.RLock()

	unsafeUnion := set.s.Union(&o.s).(*UnsafeSet)
	ret := &SafeSet{s: *unsafeUnion}
	set.RUnlock()
	o.RUnlock()
	return ret
}

func (set *SafeSet) Intersect(other Set) Set {
	o := other.(*SafeSet)

	set.RLock()
	o.RLock()

	unsafeIntersection := set.s.Intersect(&o.s).(*UnsafeSet)
	ret := &SafeSet{s: *unsafeIntersection}
	set.RUnlock()
	o.RUnlock()
	return ret
}

func (set *SafeSet) Difference(other Set) Set {
	o := other.(*SafeSet)

	set.RLock()
	o.RLock()

	unsafeDifference := set.s.Difference(&o.s).(*UnsafeSet)
	ret := &SafeSet{s: *unsafeDifference}
	set.RUnlock()
	o.RUnlock()
	return ret
}

func (set *SafeSet) SymmetricDifference(other Set) Set {
	o := other.(*SafeSet)

	set.RLock()
	o.RLock()

	unsafeDifference := set.s.SymmetricDifference(&o.s).(*UnsafeSet)
	ret := &SafeSet{s: *unsafeDifference}
	set.RUnlock()
	o.RUnlock()
	return ret
}

func (set *SafeSet) Clear() {
	set.Lock()
	set.s = newThreadUnsafeSet()
	set.Unlock()
}

func (set *SafeSet) Remove(i interface{}) {
	set.Lock()
	delete(set.s, i)
	set.Unlock()
}

func (set *SafeSet) Cardinality() int {
	set.RLock()
	defer set.RUnlock()
	return len(set.s)
}

func (set *SafeSet) Each(cb func(interface{}) bool) {
	set.RLock()
	for elem := range set.s {
		if cb(elem) {
			break
		}
	}
	set.RUnlock()
}

func (set *SafeSet) Iter() <-chan interface{} {
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

func (set *SafeSet) Iterator() *Iterator {
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

func (set *SafeSet) Equal(other Set) bool {
	o := other.(*SafeSet)

	set.RLock()
	o.RLock()

	ret := set.s.Equal(&o.s)
	set.RUnlock()
	o.RUnlock()
	return ret
}

func (set *SafeSet) Clone() Set {
	set.RLock()

	unsafeClone := set.s.Clone().(*UnsafeSet)
	ret := &SafeSet{s: *unsafeClone}
	set.RUnlock()
	return ret
}

func (set *SafeSet) String() string {
	set.RLock()
	ret := set.s.String()
	set.RUnlock()
	return ret
}

func (set *SafeSet) PowerSet() Set {
	set.RLock()
	unsafePowerSet := set.s.PowerSet().(*UnsafeSet)
	set.RUnlock()

	ret := &SafeSet{s: newThreadUnsafeSet()}
	for subset := range unsafePowerSet.Iter() {
		unsafeSubset := subset.(*UnsafeSet)
		ret.Add(&SafeSet{s: *unsafeSubset})
	}
	return ret
}

func (set *SafeSet) Pop() interface{} {
	set.Lock()
	defer set.Unlock()
	return set.s.Pop()
}

func (set *SafeSet) CartesianProduct(other Set) Set {
	o := other.(*SafeSet)

	set.RLock()
	o.RLock()

	unsafeCartProduct := set.s.CartesianProduct(&o.s).(*UnsafeSet)
	ret := &SafeSet{s: *unsafeCartProduct}
	set.RUnlock()
	o.RUnlock()
	return ret
}

func (set *SafeSet) ToSlice() []interface{} {
	keys := make([]interface{}, 0, set.Cardinality())
	set.RLock()
	for elem := range set.s {
		keys = append(keys, elem)
	}
	set.RUnlock()
	return keys
}

func (set *SafeSet) MarshalJSON() ([]byte, error) {
	set.RLock()
	b, err := set.s.MarshalJSON()
	set.RUnlock()

	return b, err
}

func (set *SafeSet) UnmarshalJSON(p []byte) error {
	set.RLock()
	err := set.s.UnmarshalJSON(p)
	set.RUnlock()

	return err
}
