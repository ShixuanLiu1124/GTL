package Set

import (
	"fmt"
)

type YourType struct {
	Name string
}

func ExampleIterator() {
	set := NewSetFromSlice([]interface{}{
		&YourType{Name: "Alise"},
		&YourType{Name: "Bob"},
		&YourType{Name: "John"},
		&YourType{Name: "Nick"},
	})

	it := set.Iterator()

	for i := range it.C {
		fmt.Println(i)
	}
}
