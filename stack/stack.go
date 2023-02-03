package stack

import "fmt"

type object interface {
	int | int32 | int64 | float32 | float64 | byte | string
}

type stack[obj object] struct {
	arr []obj
	top int64
}

//type cache[T cacheable] struct {
//	data [string]T
//}

func New[Obj object]() *stack[Obj] {
	s := stack[Obj]{}
	s.top = -1
	s.arr = make([]Obj, 0, 0)
	return &s
}

//func New[T cacheable]() cache[T] {
//	c := cache[T]{}
//	c.data = make(map[string]T)
//
//	return &c
//}
//func (c *cache[T]) Set(key string, value T) {
//	c.data[key] = value
//}

func (s *stack[obj]) Empty() bool {
	if s.top < 0 {
		return true
	}
	return false
}

//
func (s *stack[obj]) Top() obj {
	return s.arr[s.top]
}

func (s *stack[obj]) Pop() error {
	if s.Empty() {
		return fmt.Errorf(" Pop Attemted on Empty Stack")
	}
	s.top -= 1
	return nil
}

func (s *stack[obj]) Push(val obj) {
	s.arr = append(s.arr, val)
	s.top = int64(len(s.arr) - 1)
}
