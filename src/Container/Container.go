package Container

type Container interface {
	Fill() bool
	Empty() bool
	Size() int
	MaxSize() int
	SetMaxSize(int) error
	ToString() string
	Clear() bool
	CopyFromArray([]interface{}) error
}
