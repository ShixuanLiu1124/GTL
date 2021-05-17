package Container

type Container interface {
	Fill() bool
	Empty() bool
	Size() int
	MaxSize() int
	ToString() string
	Clear() bool
}
