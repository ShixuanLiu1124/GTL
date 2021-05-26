package Stack

import "fmt"

func UnsafeStackExample() {
	s, err := NewUnsafeStack(-1, 1, 2, 3, 4)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = s.Push(5)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("s.size =", s.Size())
	fmt.Println("s.Fill =", s.Fill())
	fmt.Println("s.Empty =", s.Empty())
	value, err := s.Top()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("s.Top =", value)
	fmt.Println("s.size =", s.Size())

	fmt.Println(s.String())

	s.Clear()
	fmt.Println("s.size =", s.Size())
	fmt.Println(s.String())

	s.CopyFromSlice([]interface{}{'c', "aq", []int{1, 2, 3, 4}, 1})
	fmt.Println(s.String())
}

func SafeStackExample() {

}
