package Queue

import "fmt"

func UnsafeQueueExample() {
	q, err := NewUnsafeQueue(-1, 1, 2, 3, 4, 5)
	if err != nil {
		fmt.Println(err)
	}
	err = q.Push(6)
	if err != nil {
		return
	}

	fmt.Println(q.String())
}

func SafeQueueExample() {

}
