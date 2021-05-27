package PriorityQueue

import "container/heap"

type unsafePriorityQueue struct {
	heap.Interface
}
