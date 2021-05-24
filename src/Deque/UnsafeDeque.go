package Deque

type DNode struct {
	data interface{}
	next *DNode
	prev *DNode
}

type UnsafeDeque struct {
}
