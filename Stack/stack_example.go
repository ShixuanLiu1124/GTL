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

package Stack

import "fmt"

func UnsafeStackExample() {
	s, err := NewUnsafeStack(-1, 1, 2, 3, 4, 5, 6, 7, 8)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("s.Size =", s.Size())
	v1, err := s.Top()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("s.Top =", v1)
	v1, err = s.Pop()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("v1 =", v1)
	fmt.Println("s.String =", s.String())
	b, err := s.MarshalJSON()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
	s.Clear()
	fmt.Println("s.String =", s.String())
	err = s.UnmarshalJSON(b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("s.String =", s.String())
}

func SafeStackExample() {

}
