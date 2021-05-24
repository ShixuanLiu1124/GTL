package Container

type Container interface {
	Fill() bool
	Empty() bool
	Size() int
	MaxSize() int
	SetMaxSize(int) error
	Clear() bool
	String() string
	CopyFromArray([]interface{}) error
}
