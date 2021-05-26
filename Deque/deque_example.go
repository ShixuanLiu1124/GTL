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
