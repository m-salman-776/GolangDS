package main

import "fmt"

func main() {
	st := NewStack()
	_ = st
	for i := 1; i < 10; i++ {
		st.Push(int64(i))
	}
	for st.Empty() == false {
		top := st.Top()
		fmt.Println(top)
		st.Pop()
	}
}

type Stack struct {
	arr []int64
	top int64
}

func NewStack() *Stack {
	st := &Stack{
		arr: make([]int64, 0, 0),
		top: 0,
	}
	return st
}
func (s *Stack) Empty() bool {
	if s.top < 0 {
		return true
	}
	return false
}

func (s *Stack) Top() int64 {
	return s.arr[s.top]
}
func (s *Stack) Pop() error {
	if s.Empty() {
		return fmt.Errorf(" Pop Attemted on Empty Stack")
	}
	s.top -= 1
	return nil
}
func (s *Stack) Push(val int64) {
	s.arr = append(s.arr, val)
	s.top = int64(len(s.arr) - 1)
}
