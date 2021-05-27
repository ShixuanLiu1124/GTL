package Container

type Container interface {
	Fill() bool

	Empty() bool

	Size() int

	MaxSize() int

	SetMaxSize(int) error

	Clear()

	String() string

	// CatFromSlice 从切片中复制元素到容器中
	CatFromSlice([]interface{}) error

	// ToSlice 将容器按切片形式返回
	ToSlice() []interface{}

	// MarshalJSON 将容器中的所有元素以Json数组的形式返回
	MarshalJSON() ([]byte, error)

	// UnmarshalJSON 从给定的Json数组中解析出容器,数字将被解析为json.Number
	UnmarshalJSON(b []byte) error
}
