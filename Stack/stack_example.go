package Stack

import "fmt"

func UnsafeStackExample() {
	s, err := NewUnsafeStack(-1, 1, 2, 3, 4, 5, 6, 7, 8)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("s.Size =", s.Size())
	v1, err := s.Top()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("s.Top =", v1)
	v1, err = s.Pop()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("v1 =", v1)
	fmt.Println("s.String =", s.String())
	b, err := s.MarshalJSON()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
	s.Clear()
	fmt.Println("s.String =", s.String())
	err = s.UnmarshalJSON(b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("s.String =", s.String())
}

func SafeStackExample() {

}
