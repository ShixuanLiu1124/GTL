package Set

// Iterator 用来遍历整个集合
type Iterator struct {
	// channel C用来遍历集合中的所有元素
	C <-chan interface{}
	// channel stop用来传递信号使子go程根据信号进行操作
	stop chan struct{}
}

// Stop 用于停止iterator的迭代操作， no further elements will be received on C, C will be closed.
func (i *Iterator) Stop() {
	// Allows for Stop() to be called multiple times
	// (close() panics when called on already closed channel)
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
