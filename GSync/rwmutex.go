package GSync

import (
	"context"
	"sync"
	"sync/atomic"

	"golang.org/x/sync/semaphore"
)

type RWLocker interface {
	RLock()
	RUnlock()
	WLock()
	WUnlock()
}

type DummyRWLock struct{}

func (*DummyRWLock) RLock()   {}
func (*DummyRWLock) RUnlock() {}
func (*DummyRWLock) WLock()   {}
func (*DummyRWLock) WUnlock() {}

type MutexAsRWLock struct {
	m sync.Mutex
}

func (l *MutexAsRWLock) RLock()   { l.m.Lock() }
func (l *MutexAsRWLock) RUnlock() { l.m.Unlock() }
func (l *MutexAsRWLock) WLock()   { l.m.Lock() }
func (l *MutexAsRWLock) WUnlock() { l.m.Unlock() }

type RWMutexAsRWLock struct {
	rwm sync.RWMutex
}

func (l *RWMutexAsRWLock) RLock()   { l.rwm.RLock() }
func (l *RWMutexAsRWLock) RUnlock() { l.rwm.RUnlock() }
func (l *RWMutexAsRWLock) WLock()   { l.rwm.Lock() }
func (l *RWMutexAsRWLock) WUnlock() { l.rwm.Unlock() }

// ReaderCountRWLock 最基础的读写锁
// 缺点是当读者持有锁时，写者获取锁的实现会持续自旋，
// 不断的获取锁与释放锁这一过程对CPU的计算能力来说是一种额外的消耗。
type ReaderCountRWLock struct {
	// m 互斥锁
	m sync.Mutex
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

type ReaderCountCondRWLock struct {
	readerCount int
	c           *sync.Cond
}

// NewReaderCountCondRWLock creates a new ReaderCountRWLock.
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
	} else if l.readerCount == 0 {
		l.c.Signal()
	}
	l.c.L.Unlock()
}

func (l *ReaderCountCondRWLock) WLock() {
	l.c.L.Lock()
	for l.readerCount > 0 {
		l.c.Wait()
	}
}

func (l *ReaderCountCondRWLock) WUnlock() {
	l.c.Signal()
	l.c.L.Unlock()
}

const maxWeight int64 = 1 << 30

type SemaRWLock struct {
	s *semaphore.Weighted
}

// NewSemaRWLock creates a new SemaRWLock.
func NewSemaRWLock() *SemaRWLock {
	return &SemaRWLock{semaphore.NewWeighted(maxWeight)}
}

func (l *SemaRWLock) RLock() {
	l.s.Acquire(context.Background(), 1)
}

func (l *SemaRWLock) RUnlock() {
	l.s.Release(1)
}

func (l *SemaRWLock) WLock() {
	l.s.Acquire(context.Background(), maxWeight)
}

func (l *SemaRWLock) WUnlock() {
	l.s.Release(maxWeight)
}

type WritePreferRWLock struct {
	readerCount int
	hasWriter   bool
	c           *sync.Cond
}

func NewWritePreferRWLock() *WritePreferRWLock {
	return &WritePreferRWLock{0, false, sync.NewCond(new(sync.Mutex))}
}

func (l *WritePreferRWLock) RLock() {
	l.c.L.Lock()

	for l.hasWriter {
		l.c.Wait()
	}
	l.readerCount++
	l.c.L.Unlock()
}

func (l *WritePreferRWLock) RUnlock() {
	l.c.L.Lock()
	l.readerCount--
	if l.readerCount == 0 {
		l.c.Broadcast()
	}
	l.c.L.Unlock()
}

func (l *WritePreferRWLock) WLock() {
	l.c.L.Lock()

	for l.hasWriter {
		l.c.Wait()
	}
	l.hasWriter = true
	for l.readerCount > 0 {
		l.c.Wait()
	}
	l.c.L.Unlock()
}

func (l *WritePreferRWLock) WUnlock() {
	l.c.L.Lock()
	l.hasWriter = false
	l.c.Broadcast()
	l.c.L.Unlock()
}

type WritePreferFastRWLock struct {
	w sync.Mutex

	writerWait chan struct{}
	readerWait chan struct{}

	numPending int32

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
	if atomic.AddInt32(&l.numPending, 1) < 0 {
		<-l.readerWait
	}
}

func (l *WritePreferFastRWLock) RUnlock() {
	if r := atomic.AddInt32(&l.numPending, -1); r < 0 {
		if r+1 == 0 || r+1 == -maxReaders {
			panic("RUnlock of unlocked RWLock")
		}
		if atomic.AddInt32(&l.readersDeparting, -1) == 0 {
			l.writerWait <- struct{}{}
		}
	}
}

func (l *WritePreferFastRWLock) WLock() {
	l.w.Lock()
	r := atomic.AddInt32(&l.numPending, -maxReaders) + maxReaders
	if r != 0 && atomic.AddInt32(&l.readersDeparting, r) != 0 {
		<-l.writerWait
	}
}

func (l *WritePreferFastRWLock) WUnlock() {
	r := atomic.AddInt32(&l.numPending, maxReaders)
	if r >= maxReaders {
		panic("WUnlock of unlocked RWLock")
	}
	for i := 0; i < int(r); i++ {
		l.readerWait <- struct{}{}
	}
	l.w.Unlock()
}
