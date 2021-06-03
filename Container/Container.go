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

package Container

type Container interface {
	Fill() bool

	Empty() bool

	Size() int

	MaxSize() int

	SetMaxSize(maxSize int) error

	Clear()

	String() string

	// CatFromSlice 从切片中复制元素到容器中
	CatFromSlice(values []interface{}) error

	// ToSlice 将容器按切片形式返回
	ToSlice() []interface{}

	// MarshalJSON 将容器中的所有元素以Json数组的形式返回
	MarshalJSON() ([]byte, error)

	// UnmarshalJSON 从给定的Json数组中解析出容器,数字将被解析为json.Number
	UnmarshalJSON(b []byte) error
}
