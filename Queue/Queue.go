package Queue

type Queue interface {
	Push(value interface{}) error

	Front() (interface{}, error)

	Pop() (interface{}, error)
}
