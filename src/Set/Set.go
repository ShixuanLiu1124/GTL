package Set

type Set interface {
	Add(i interface{}) bool
	Cardinality() int
	Clear()
	Clone() Set
	Contains(i ...interface{}) bool
	Difference(other Set) Set
	Equal(other Set) bool
	Intersect(other Set) Set
	IsProperSubset(other Set) bool
	IsProperSuperset(other Set) bool
	IsSubset(other Set) bool
	IsSuperset(other Set) bool
	Each(func(interface{}) bool)
	Iter() <-chan interface{}
	Iterator() *iterator
	Remove(i interface{})
	String() string
	SymmetricDifference(other Set) Set
	Union(other Set) Set
	Pop() interface{}
	PowerSet() Set
	CartesianProduct(other Set) Set
	ToSlice() []interface{}
}

// NewSet creates and returns a reference to an empty set.  Operations
// on the resulting set are thread-safe.
func NewSet(s ...interface{}) Set {
	set := newThreadSafeSet()
	for _, item := range s {
		set.Add(item)
	}
	return &set
}

// NewSetWith creates and returns a new set with the given elements.
// Operations on the resulting set are thread-safe.
func NewSetWith(elts ...interface{}) Set {
	return NewSetFromSlice(elts)
}

// NewSetFromSlice creates and returns a reference to a set from an
// existing slice.  Operations on the resulting set are thread-safe.
func NewSetFromSlice(s []interface{}) Set {
	a := NewSet(s...)
	return a
}

// NewThreadUnsafeSet creates and returns a reference to an empty set.
// Operations on the resulting set are not thread-safe.
func NewThreadUnsafeSet() Set {
	set := newThreadUnsafeSet()
	return &set
}

// NewThreadUnsafeSetFromSlice creates and returns a reference to a
// set from an existing slice.  Operations on the resulting set are
// not thread-safe.
func NewThreadUnsafeSetFromSlice(s []interface{}) Set {
	a := NewThreadUnsafeSet()
	for _, item := range s {
		a.Add(item)
	}
	return a
}
