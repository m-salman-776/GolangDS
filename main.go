package main

import (
	"fmt"
	"github.com/malviyansalman/GolangDS/linkedlist"
)

func main() {
	//st := stack.New[byte]()
	//for _, i := range []byte{'a', 'b', 'c', 'd'} {
	//	//s := fmt.Sprintf("%v", i)
	//	st.Push(i)
	//}
	//for st.Empty() == false {
	//	top := st.Top()
	//	fmt.Println(top)
	//	st.Pop()
	//}

	ll := linkedlist.NewDoublyLinkedList[string]()
	for i := 1; i < 4; i++ {
		//ll = ll.AddFront(i)
		f := fmt.Sprintf("%v_k", i)
		ll.AddLast(f)
	}
	temp := ll
	for ll != nil {
		fmt.Println(ll.Val)
		temp = ll
		ll = ll.Next
		//if ll.Next != nil{

		//}
	}
	for temp != nil {
		fmt.Println(temp.Val)
		temp = temp.Previous
	}

}
