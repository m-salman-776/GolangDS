package queue

import "container/list"

type object interface {
	int | int32 | int64 | string | byte
}

type queue[o object] struct {
	list list.List
}

// todo
