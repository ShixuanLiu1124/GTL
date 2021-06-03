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

package Deque

import "fmt"

func UnsafeQueueExample() {
	q, err := NewUnsafeDeque(-1, 1, 2, 3, 4, 5)
	if err != nil {
		fmt.Println(err)
	}
	err = q.PushBack(6)
	if err != nil {
		return
	}

	fmt.Println(q.String())
	v1, err := q.Front()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("v1 =", v1)
	v2, err := q.PopBack()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("v2 =", v2)
	fmt.Println("q.String =", q.String())
	fmt.Println("q.Size =", q.Size())

	b, err := q.MarshalJSON()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("b =", string(b))
	q.Clear()
	fmt.Println("q.String =", q.String())
	err = q.UnmarshalJSON(b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("q.String =", q.String())
}

func SafeQueueExample() {

}
