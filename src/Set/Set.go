package Set

type Set interface {
	Add(i interface{}) bool
	Size() int
	Clear()
	Clone() Set
	Contains(...interface{}) bool

	// Difference 求s - other差集
	Difference(other Set) Set

	// Equal 判断两个集合是否相等
	Equal(other Set) bool

	// Intersect 求该集合s和other的交集
	Intersect(other Set) Set

	// Union 求该集合s和other的并集
	Union(other Set) Set

	// IsProperSubset 判断other是否是该集合s的真子集
	IsProperSubset(other Set) bool

	// IsProperSuperset 判断other是否是该集合s的真超集
	IsProperSuperset(other Set) bool

	// IsSubset 判断other是否是该集合s的子集
	IsSubset(other Set) bool

	// IsSuperset 判断other是否是该集合s的超集
	IsSuperset(other Set) bool

	// Iter 返回一个可以遍历该集合s的通道
	Iter() <-chan interface{}

	// Iterator 返回该集合s的一个迭代器
	Iterator() *Iterator
	Remove(i interface{})

	// SymmetricDifference 求该集合s和other的对称差集
	// 对称差集：只属于其中一个集合，而不属于另一个集合的元素组成的集合。
	SymmetricDifference(other Set) Set

	// CartesianProduct 求该集合s和other的笛卡尔积
	CartesianProduct(other Set) Set
	ToSlice() []interface{}
	String() string
}

// NewSet NewSet创建并返回一个线程安全的空集的引用
func NewSet(s ...interface{}) Set {
	set := newSafeSet()
	for _, item := range s {
		set.Add(item)
	}
	return &set
}

// NewSetWith 依照给定的元素创建一个线程安全的空集的引用
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
	set := newUnsafeSet()
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
