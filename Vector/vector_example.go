package Vector

import (
	"fmt"
)

func UnsafeVectorExample() {
	// 对用离散元素创建Vector进行设置
	v, err := NewUnsafeVector(-1, 2, 3, 4, 5, 6)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("v.Size =", v.Size())
	fmt.Println("v.String =", v.String())
	fmt.Println("v.maxSize =", v.maxSize)
	for i := 0; i < v.Size(); i++ {
		value, err := v.At(i)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%v ", value)
	}
	fmt.Printf("\n")
	err = v.SetMaxSize(20)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("v.maxSize =", v.MaxSize())
	p := v.Find(3, func(i interface{}, i2 interface{}) bool {
		v1, _ := i.(int)
		v2, _ := i2.(int)
		return v1 < v2
	})
	fmt.Println("v.Find =", p)

	b, err := v.MarshalJSON()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("b =", string(b))
	v.Clear()
	fmt.Println("v.String =", v.String())
	v.UnmarshalJSON(b)
	fmt.Println("v.String =", v.String())
}
