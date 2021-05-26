package Container

type Container interface {
	Fill() bool
	Empty() bool
	Size() int
	MaxSize() int
	SetMaxSize(int) error
	Clear()
	String() string
	CopyFromArray([]interface{}) error
	ToSlice() []interface{}
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(b []byte) error
}
