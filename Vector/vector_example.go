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

package Vector

import (
	"fmt"
)

func UnsafeVectorExample() {
	// 对用离散元素创建Vector进行设置
	v, err := NewUnsafeVector(-1, 2, 3, 4, 5, 6)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("v.Size =", v.Size())
	fmt.Println("v.String =", v.String())
	fmt.Println("v.maxSize =", v.maxSize)
	for i := 0; i < v.Size(); i++ {
		value, err := v.At(i)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%v ", value)
	}
	fmt.Printf("\n")
	err = v.SetMaxSize(20)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("v.maxSize =", v.MaxSize())
	p := v.Find(3, func(i interface{}, i2 interface{}) bool {
		v1, _ := i.(int)
		v2, _ := i2.(int)
		return v1 < v2
	})
	fmt.Println("v.Find =", p)

	b, err := v.MarshalJSON()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("b =", string(b))
	v.Clear()
	fmt.Println("v.String =", v.String())
	v.UnmarshalJSON(b)
	fmt.Println("v.String =", v.String())
}
