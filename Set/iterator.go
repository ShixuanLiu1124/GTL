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

// Iterator 用来遍历整个集合
type Iterator struct {
	// channel C用来遍历集合中的所有元素
	C <-chan interface{}
	// channel stop用来传递信号使子go程根据信号进行操作
	stop chan struct{}
}

// Stop 用于停止iterator的迭代操作，当C不再接收元素时，C会被关闭
func (i *Iterator) Stop() {
	// Stop能被多次调用
	defer func() {
		recover()
	}()

	close(i.stop)

	// 消除C中剩下的元素
	for range i.C {
	}
}

// newIterator 返回一个迭代器、迭代器中的C和stopChan
func newIterator() (*Iterator, chan<- interface{}, <-chan struct{}) {
	itemChan := make(chan interface{})
	stopChan := make(chan struct{})
	return &Iterator{
		C:    itemChan,
		stop: stopChan,
	}, itemChan, stopChan
}
