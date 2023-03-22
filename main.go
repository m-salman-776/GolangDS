package main

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

	//ll := linkedlist.NewDoublyLinkedList[string]()
	//for i := 1; i < 4; i++ {
	//	//ll = ll.AddFront(i)
	//	f := fmt.Sprintf("%v_k", i)
	//	ll.AddLast(f)
	//}
	//temp := ll
	//for ll != nil {
	//	fmt.Println(ll.Val)
	//	temp = ll
	//	ll = ll.Next
	//	//if ll.Next != nil{
	//
	//	//}
	//}
	//for temp != nil {
	//	fmt.Println(temp.Val)
	//	temp = temp.Previous
	//}
	//ss := BreakPalindrome("abccba")
	//_ = ss
	//fmt.Println("DNE")

	//q := queue.New[int]()
	//for i := 0; i < 10; i++ {
	//	q.Push(i)
	//}
	//for q.Empty() == false {
	//	node := q.Front()
	//	fmt.Println(node)
	//	q.Pop()
	//}
}
func BreakPalindrome(strs string) string {
	str := []rune(strs)
	for {
		change := false
		for i := 0; i < len(str); i++ {
			if str[i] != 'a' {
				str[i] = 'a'
				change = true
				break
			}
		}
		if !change {
			str[len(str)-1] = 'b'
		}
		if palindrome(str) == false {
			return string(str)
		}
	}
	return ""
}

func palindrome(str []rune) bool {
	left := 0
	right := len(str) - 1
	for left < right {
		if str[left] != str[right] {
			return false
		}
		left += 1
		right -= 1
	}
	return true
}
