package main

import (
	"GolangDS/stack"
	"fmt"
)

func main() {
	st := stack.New[byte]()
	for _, i := range []byte{'a', 'b', 'c', 'd'} {
		//s := fmt.Sprintf("%v", i)
		st.Push(i)
	}
	for st.Empty() == false {
		top := st.Top()
		fmt.Println(top)
		st.Pop()
	}
}
