package queue

type object interface {
	int | int32 | int64 | string | byte
}
type node[o object] struct {
	data o
	next *node[o]
}

func makeNode[o object](data o) *node[o] {
	n := &node[o]{data: data, next: nil}
	return n
}

type queue[o object] struct {
	size        int64
	front, back *node[o]
}

func New[o object]() *queue[o] {
	q := &queue[o]{}
	return q
}
func (q *queue[o]) Size(data o) int64 {
	return q.size
}

func (q *queue[o]) Push(data o) {
	node := makeNode(data)
	if q.size == 0 {
		q.front = node
		q.back = node
	} else {
		q.back.next = node
		q.back = q.back.next
	}
	q.size += 1
}

func (q *queue[o]) Empty() bool {
	if q.size > 0 {
		return false
	}
	return true
}

func (q *queue[o]) Front() o {
	return q.front.data
}
func (q *queue[o]) Pop() {
	q.front = q.front.next
	q.size -= 1
}

// todo
