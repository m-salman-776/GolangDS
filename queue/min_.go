package queue

import "sync"

type minHeap[o object] struct {
	arr            []o
	size           int64
	minRebalanceTh int64
	maxRebalanceTh int64
	cmp            func(a o, b o)
	wh             sync.WaitGroup
	lock           sync.Mutex
}

type maxHeap[o object] struct {
	arr          []o
	size         int64
	minRebalance int64
	maxRebalance int64
	cmp          func(a o, b o)
}

func NewMinHeap[o object]() *minHeap[o] {
	mh := &minHeap[o]{}
	return mh
}

func (m *minHeap[o]) Empty() bool {
	m.wh.Wait()
	m.lock.Lock()
	if m.getSize() > 0 {
		return false
	}
	return true
}
func (m *minHeap[o]) Push(data o) {
	m.arr[m.size] = data
	m.balance(m.size)
	m.size += 1
}
func (m *minHeap[o]) Pop() {
	if m.Empty() {
		panic("Pop on Empty")
	}
	m.swap(0, m.size-1)
	m.heapify(0)
	m.size -= 1
}
func (m *minHeap[o]) Top() o {
	return m.arr[0]
}
func (m *minHeap[o]) heapify(idx int64) {
	leftChild := m.leftChile(idx)
	rightChild := m.rightChild(idx)
	smallest := idx
	if leftChild < m.getSize() && m.arr[leftChild] < m.arr[idx] {
		smallest = leftChild
	}
	if rightChild < m.getSize() && m.arr[rightChild] < m.arr[idx] {
		smallest = rightChild
	}
	if smallest != idx {
		m.arr[smallest], m.arr[idx] = m.arr[smallest], m.arr[idx]
		m.heapify(smallest)
	}
}

func (m *minHeap[o]) balance(lastIdx int64) {
	for lastIdx >= 0 && m.arr[m.parent(lastIdx)] > m.arr[lastIdx] {
		parentIdx := m.parent(lastIdx)
		m.arr[parentIdx], m.arr[lastIdx] = m.arr[lastIdx], m.arr[parentIdx]
		lastIdx = parentIdx
	}
}

func (m *minHeap[o]) swap(idx1, idx2 int64) {
	m.arr[idx1], m.arr[idx2] = m.arr[idx2], m.arr[idx1]
}

func (m *minHeap[o]) getSize() int64 {
	return m.size + 1
}

func (m *minHeap[o]) parent(idx int64) int64 {
	return idx / 2
}
func (m *minHeap[o]) leftChile(idx int64) int64 {
	return 2*idx + 1
}
func (m *minHeap[o]) rightChild(idx int64) int64 {
	return 2*idx + 2
}
