package linkedlist

type object interface {
	int | int32 | int64 | string | byte
}

type linkedList[o object] struct {
	Val   o
	Next  *linkedList[o]
	valid bool // to check if we really have first element
}
type doublyLinkedList[o object] struct {
	Val      o
	Next     *doublyLinkedList[o]
	Previous *doublyLinkedList[o]
	valid    bool
}

func NewLinkedList[Obj object]() *linkedList[Obj] {
	ll := &linkedList[Obj]{}
	return ll
}
func NewDoublyLinkedList[Obj object]() *doublyLinkedList[Obj] {
	dll := &doublyLinkedList[Obj]{}
	dll.valid = false
	return dll
}
func makeLlNode[o object](val o) *linkedList[o] {
	ll := &linkedList[o]{
		Val:   val,
		Next:  nil,
		valid: true,
	}
	return ll
}
func (l *linkedList[o]) AddLast(val o) *linkedList[o] {
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
func (l *linkedList[o]) AddFront(val o) *linkedList[o] {
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
