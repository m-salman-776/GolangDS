package linkedlist

type object interface {
	int | int32 | int64 | string | byte
}

type LinkedList[o object] struct {
	Val   o
	Next  *LinkedList[o]
	valid bool // to check if we really have first element
}
type doublyLinkedList[o object] struct {
	Val      o
	Next     *doublyLinkedList[o]
	Previous *doublyLinkedList[o]
	valid    bool
}

func NewLinkedList[Obj object]() *LinkedList[Obj] {
	ll := &LinkedList[Obj]{}
	return ll
}
func NewDoublyLinkedList[Obj object]() *doublyLinkedList[Obj] {
	dll := &doublyLinkedList[Obj]{}
	dll.valid = false
	return dll
}
func makeLlNode[o object](val o) *LinkedList[o] {
	ll := &LinkedList[o]{
		Val:   val,
		Next:  nil,
		valid: true,
	}
	return ll
}
func (l *LinkedList[o]) AddLast(val o) *LinkedList[o] {
	ll := makeLlNode(val)
	if l.valid == false {
		return ll
	}
	temp := l
	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = ll
	return l
}
func (l *LinkedList[o]) AddFront(val o) *LinkedList[o] {
	ll := makeLlNode(val)
	if l.valid == false {
		return ll
	}
	ll.Next = l
	return ll
}
func makeDllNode[o object](val o) *doublyLinkedList[o] {
	ll := &doublyLinkedList[o]{
		Val:      val,
		Next:     nil,
		Previous: nil,
		valid:    true,
	}
	return ll
}
func (l *doublyLinkedList[o]) AddLast(val o) *doublyLinkedList[o] {
	ll := makeDllNode(val)
	if l.valid == false {
		l.Val = val
		l.valid = true
		return l
	}
	temp := l
	for temp.Next != nil {
		temp = temp.Next
	}
	ll.Previous = temp
	temp.Next = ll
	return l
}
func (l *doublyLinkedList[o]) AddFront(val o) *doublyLinkedList[o] {
	ll := makeDllNode(val)
	if l.valid == false {
		l.Val = val
		l.valid = true
		return l
	}
	ll.Next = l
	l.Previous = ll
	return ll
}
