package Stack

import "sync"

type safeStack struct {
	us *unsafeStack
	m  *sync.RWMutex
}

func (s *safeStack) Push(value interface{}) error {
	s.m.Lock()
	defer s.m.Unlock()

	return s.us.Push(value)
}

func (s *safeStack) Top() (interface{}, error) {
	s.m.RLock()
	defer s.m.RUnlock()

	return s.us.Top()
}

func (s *safeStack) Pop() (interface{}, error) {
	s.m.Lock()
	defer s.m.Unlock()

	return s.us.Pop()
}

func (s *safeStack) Fill() bool {
	s.m.RLock()
	defer s.m.RUnlock()

	return s.us.Fill()
}

func (s *safeStack) Empty() bool {
	s.m.RLock()
	defer s.m.RUnlock()

	return s.us.Empty()
}

func (s *safeStack) Size() int {
	s.m.RLock()
	defer s.m.RUnlock()

	return s.us.Size()
}

func (s *safeStack) MaxSize() int {
	s.m.RLock()
	defer s.m.RUnlock()

	return s.us.MaxSize()
}

func (s *safeStack) SetMaxSize(maxSize int) error {
	s.m.Lock()
	defer s.m.Unlock()

	return s.us.SetMaxSize(maxSize)
}

func (s *safeStack) Clear() {
	s.m.Lock()
	defer s.m.Unlock()

	s.us.Clear()
}

func (s *safeStack) String() string {
	s.m.RLock()
	defer s.m.RUnlock()

	return s.us.String()
}

func (s *safeStack) CatFromSlice(values []interface{}) error {
	s.m.Lock()
	defer s.m.Unlock()

	return s.us.CatFromSlice(values)
}

func (s *safeStack) ToSlice() []interface{} {
	s.m.Lock()
	defer s.m.Unlock()

	return s.us.ToSlice()
}

func (s *safeStack) MarshalJSON() ([]byte, error) {
	s.m.RLock()
	defer s.m.RUnlock()

	return s.us.MarshalJSON()
}

func (s *safeStack) UnmarshalJSON(b []byte) error {
	s.m.Lock()
	defer s.m.Unlock()

	return s.us.UnmarshalJSON(b)
}
