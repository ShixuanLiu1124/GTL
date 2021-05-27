package Stack

type Stack interface {
	Push(value interface{}) error

	Top() (interface{}, error)

	Pop() (interface{}, error)
}
