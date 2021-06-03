package GSync

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"

	"golang.org/x/sync/semaphore"
)

type RWLocker interface {
	// RLock 读者读数据时获取读锁
	RLock()

	// RUnlock 读者读完数据时释放读锁
	RUnlock()

	// WLock 写者写数据时获取写锁
	WLock()

	// WUnlock 写者写完数据时释放写锁
	WUnlock()
}

// ReaderCountRWLock 最基础的读写锁
// 缺点是当读者持有锁时，写者获取锁的实现会持续自旋
// 不断的获取锁与释放锁这一过程对CPU的计算能力来说是一种额外的消耗
type ReaderCountRWLock struct {
	// m 互斥锁
	m *sync.Mutex

	// readerCount 记录读者数量
	readerCount int
}

// RLock 读进入的时候对readerCount变量进行访问控制
func (l *ReaderCountRWLock) RLock() {
	// 读的时候给readerCount上锁
	l.m.Lock()

	// 读者数量加一
	l.readerCount++

	// 给readerCount解锁
	l.m.Unlock()
}

// RUnlock 读退出的时候对readerCount变量进行访问控制
func (l *ReaderCountRWLock) RUnlock() {
	l.m.Lock()
	l.readerCount--
	l.m.Unlock()
}

// WLock 写者对mutex上锁，同时检查是否存在读者持有锁，如果存在，写者释放mutex并再次尝试，这叫自旋。
func (l *ReaderCountRWLock) WLock() {
	for {
		l.m.Lock()
		if l.readerCount > 0 {
			l.m.Unlock()
		} else {
			break
		}
	}
}

// WUnlock 读者释放锁
func (l *ReaderCountRWLock) WUnlock() {
	l.m.Unlock()
}

// ReaderCountCondRWLock 使用条件变量sync.Cond避免自旋操作，实现高效等待
type ReaderCountCondRWLock struct {
	readerCount int
	c           *sync.Cond
}

// NewReaderCountCondRWLock 返回一个新的ReaderCountRWLock指针
func NewReaderCountCondRWLock() *ReaderCountCondRWLock {
	return &ReaderCountCondRWLock{0, sync.NewCond(new(sync.Mutex))}
}

func (l *ReaderCountCondRWLock) RLock() {
	l.c.L.Lock()
	l.readerCount++
	l.c.L.Unlock()
}

func (l *ReaderCountCondRWLock) RUnlock() {
	l.c.L.Lock()
	l.readerCount--
	if l.readerCount < 0 {
		panic("readerCount negative")
	} else if l.readerCount == 0 { // 没有读者时，在条件变量上发出信号以唤起一个等待线程
		l.c.Signal()
	}
	l.c.L.Unlock()
}

func (l *ReaderCountCondRWLock) WLock() {
	l.c.L.Lock()
	// 还有读者时
	for l.readerCount > 0 {
		// Wait过程仍然处于循环中，因为极有可能在读者发出信号之后、写者获取锁之前，另一个读者先拿到锁。
		l.c.Wait()
	}
}

func (l *ReaderCountCondRWLock) WUnlock() {
	// 唤醒一个等待的线程
	l.c.Signal()

	// 写者释放锁
	l.c.L.Unlock()
}

// SemaRWLock 使用golang.org/x/sync/semaphore包实现
// 使用带权重的信号量类型semaphore.Weighted
type SemaRWLock struct {
	s *semaphore.Weighted
}

// 设定最大权重maxWeight
const maxWeight int64 = 1 << 30

// NewSemaRWLock 返回一个新的SemaRWLock指针
func NewSemaRWLock() *SemaRWLock {
	return &SemaRWLock{semaphore.NewWeighted(maxWeight)}
}

func (l *SemaRWLock) RLock() {
	// 阻塞的获取指定权重的资源
	// 因为可以同时有多个读者线程进行读操作，所以每个读者线程就获取1权重的资源
	err := l.s.Acquire(context.Background(), 1)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (l *SemaRWLock) RUnlock() {
	l.s.Release(1)
}

func (l *SemaRWLock) WLock() {
	// 因为同一时刻只能有一个写者线程进行写操作，所以每个写者线程获取maxWeight权重的资源
	err := l.s.Acquire(context.Background(), maxWeight)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (l *SemaRWLock) WUnlock() {
	l.s.Release(maxWeight)
}

/*
上述的三种实现都存在一个问题：当读者数量很大时，可能会导致写者饥饿。
例如，第一版实现中 readerCount 为0时，写者才能够获取锁，
假设有两个活跃的读者以及一个等待的写者，在写者等待一个读者释放锁的过程中，另一个读者可能又会获取锁
*/

// WritePreferRWLock 写者优先锁
type WritePreferRWLock struct {
	readerCount int

	// hasWriter 当有写者等待获取锁的时候，值为true
	hasWriter bool
	c         *sync.Cond
}

// NewWritePreferRWLock 返回一个新的WritePreferRWLock指针
func NewWritePreferRWLock() *WritePreferRWLock {
	return &WritePreferRWLock{0, false, sync.NewCond(new(sync.Mutex))}
}

func (l *WritePreferRWLock) RLock() {
	// 读者获取锁
	l.c.L.Lock()

	// 检查是否有写者等待获取锁
	// 如果有的话将会让出获取锁的权限，保证写者先获取锁。
	for l.hasWriter {
		l.c.Wait()
	}

	// 读者自己获取锁
	l.readerCount++
	l.c.L.Unlock()
}

func (l *WritePreferRWLock) RUnlock() {
	l.c.L.Lock()
	l.readerCount--
	if l.readerCount == 0 {
		// 唤醒所有线程，加速让写者获得锁
		l.c.Broadcast()
	}
	l.c.L.Unlock()
}

// WLock 写者在WLock与WUnlock之间不再持有mutex，取而代之，mutex仅用于控制共享结构的访问
func (l *WritePreferRWLock) WLock() {
	l.c.L.Lock()

	for l.hasWriter {
		l.c.Wait()
	}
	l.hasWriter = true
	for l.readerCount > 0 {
		l.c.Wait()
	}

	// 因为有hasWriter变量为true保证其他线程都被阻塞，所以此时可以释放锁
	l.c.L.Unlock()
}

func (l *WritePreferRWLock) WUnlock() {
	l.c.L.Lock()
	l.hasWriter = false

	// 使用Broadcast而不是Signal是因为可能存在多个读者等待，而我们期望唤醒所有等待的读者。
	l.c.Broadcast()
	l.c.L.Unlock()
}

// WritePreferFastRWLock 模仿Go语言自身RWMutex实现更高效的读写锁
type WritePreferFastRWLock struct {
	w *sync.Mutex

	writerWait chan struct{}
	readerWait chan struct{}

	// numPending 已经持有锁的读者数量
	// 写者将该属性减去maxReaders，如果得到一个负数就表明一个写者正在使用锁
	numPending int32

	// readersDeparting 在写者持有锁之前获取锁的读者数量（读者释放锁，也会随之减1）
	readersDeparting int32
}

const maxReaders int32 = 1 << 30

func NewWritePreferFastRWLock() *WritePreferFastRWLock {
	var l WritePreferFastRWLock
	l.writerWait = make(chan struct{})
	l.readerWait = make(chan struct{})
	return &l
}

func (l *WritePreferFastRWLock) RLock() {
	// 通过原子操作atomic包中的操作访问该字段，因此不再需要锁
	// 如果numPending是非负数，表明没有写者等待持有锁或正持有锁，因此读者可以继续操作
	// 如果numPending是负数，表明写者正在等待获取锁或已经获取锁，因此读者将会让出权限，保持等待
	if atomic.AddInt32(&l.numPending, 1) < 0 {
		// 保持等待是通过在一个无缓冲channel上等待实现的。
		<-l.readerWait
	}
}

func (l *WritePreferFastRWLock) RUnlock() {
	// 读者释放锁时，将numPending减1
	if r := atomic.AddInt32(&l.numPending, -1); r < 0 {
		// unlock情况下的numPending应该是等于0的，此时再对numPending进行减1的话就会让它变为-1，这是非法情况
		// 或者是在写者已经占用临界区的时候调用RUnlock会导致numPending = -maxReaders - 1，这也是非法情况
		if r+1 == 0 || r+1 == -maxReaders {
			panic("RUnlock of unlocked RWLock")
		}

		// 如果等于0说明已经没有读者在临界区
		if atomic.AddInt32(&l.readersDeparting, -1) == 0 {
			// 放入空结构体唤醒写者线程
			l.writerWait <- struct{}{}
		}
	}
}

func (l *WritePreferFastRWLock) WLock() {
	l.w.Lock()
	// 通过执行numPending减maxReaders操作来告知读者有一个写者在申请临界区
	r := atomic.AddInt32(&l.numPending, -maxReaders) + maxReaders

	// r!=0表示有读者在临界区内，此时将r加到readersDeparting中，让读者知道有多少个尝试持有锁或持有锁的读者
	if r != 0 && atomic.AddInt32(&l.readersDeparting, r) != 0 {
		<-l.writerWait
	}
}

func (l *WritePreferFastRWLock) WUnlock() {
	// 告知读者，写者已经占用完了临界区
	r := atomic.AddInt32(&l.numPending, maxReaders)

	// 如果本来就没有写锁，便报出panic
	if r >= maxReaders {
		panic("WUnlock of unlocked RWLock")
	}
	// 通知所有读者（数量为r），写者已经占用完了临界区
	for i := 0; i < int(r); i++ {
		l.readerWait <- struct{}{}
	}
	l.w.Unlock()
}
