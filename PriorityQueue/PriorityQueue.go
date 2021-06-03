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

package PriorityQueue

import (
	"GTL/Container"
)

type PriorityQueue interface {
	Push(values interface{}) error

	Pop() (interface{}, error)

	Top() (interface{}, error)

	swap(i, j int)

	up(index int)

	down(start int, end int) bool

	fix(index int)

	SetFunc(less func(interface{}, interface{}) bool)

	Container.Container
}
