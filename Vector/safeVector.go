package Vector

import "sync"

type safeVector struct {
	uv *unsafeVector
	m  *sync.RWMutex
}

func (v *safeVector) PushBack(value interface{}) error {
	v.m.Lock()
	defer v.m.Unlock()

	return v.uv.PushBack(value)
}

func (v *safeVector) PopBack() (interface{}, error) {
	v.m.Lock()
	defer v.m.Unlock()

	return v.uv.PopBack()
}

func (v *safeVector) Set(index, value int) error {
	v.m.Lock()
	defer v.m.Unlock()

	return v.uv.Set(index, value)
}

func (v *safeVector) At(index int) (interface{}, error) {
	v.m.RLock()
	defer v.m.RUnlock()

	return v.uv.At(index)
}

func (v *safeVector) Remove(start, end int) error {
	v.m.Lock()
	defer v.m.Unlock()

	return v.uv.Remove(start, end)
}

func (v *safeVector) Find(value interface{}, less func(interface{}, interface{}) bool) int {
	v.m.RLock()
	defer v.m.RUnlock()

	return v.uv.Find(value, less)
}

func (v *safeVector) Fill() bool {
	v.m.RLock()
	defer v.m.RUnlock()

	return v.uv.Fill()
}

func (v *safeVector) Empty() bool {
	v.m.RLock()
	defer v.m.RUnlock()

	return v.uv.Empty()
}

func (v *safeVector) Size() int {
	v.m.RLock()
	defer v.m.RUnlock()

	return v.uv.Size()
}

func (v *safeVector) MaxSize() int {
	v.m.RLock()
	defer v.m.RUnlock()

	return v.uv.MaxSize()
}

func (v *safeVector) SetMaxSize(maxSize int) error {
	v.m.Lock()
	defer v.m.Unlock()

	return v.uv.SetMaxSize(maxSize)
}

func (v *safeVector) Clear() {
	v.m.Lock()
	defer v.m.Unlock()

	v.uv.Clear()
}

func (v *safeVector) String() string {
	v.m.RLock()
	defer v.m.RUnlock()

	return v.uv.String()
}

func (v *safeVector) CatFromSlice(values []interface{}) error {
	v.m.Lock()
	defer v.m.Unlock()

	return v.uv.CatFromSlice(values)
}

func (v *safeVector) ToSlice() []interface{} {
	v.m.RLock()
	defer v.m.RUnlock()

	return v.uv.ToSlice()
}

func (v *safeVector) MarshalJSON() ([]byte, error) {
	v.m.RLock()
	defer v.m.RUnlock()

	return v.uv.MarshalJSON()
}

func (v *safeVector) UnmarshalJSON(b []byte) error {
	v.m.Lock()
	defer v.m.Unlock()

	return v.uv.UnmarshalJSON(b)
}
