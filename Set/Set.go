/*
 *  Copyright (C) 2021  Shixuan Liu
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU General Public License for more details.
 *
 *     You should have received a copy of the GNU General Public License
 *     along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package Set

import "GTL/Container"

type Set interface {
	Insert(value interface{}) error
	Clone() Set
	Contains(values ...interface{}) bool

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
	Remove(value interface{})

	// SymmetricDifference 求该集合s和other的对称差集
	// 对称差集：只属于其中一个集合，而不属于另一个集合的元素组成的集合。
	SymmetricDifference(other Set) Set

	// CartesianProduct 求该集合s和other的笛卡尔积
	CartesianProduct(other Set) Set

	Container.Container
}
