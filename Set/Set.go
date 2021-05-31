package Set

import "GTL/Container"

type Set interface {
	Insert(i interface{}) error
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

	Container.Container
}
