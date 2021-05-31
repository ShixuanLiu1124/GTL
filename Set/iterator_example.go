package Set

import (
	"fmt"
)

type YourType struct {
	Name string
}

func ExampleIterator() {
	set, err := NewSafeSetWithSlice(-1, []interface{}{
		&YourType{Name: "Alise"},
		&YourType{Name: "Bob"},
		&YourType{Name: "John"},
		&YourType{Name: "Nick"},
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	it := set.Iterator()

	for i := range it.C {
		fmt.Println(i)
	}
}
