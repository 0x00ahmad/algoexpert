package main

import "math"

// Do not edit the class below except for the buildHeap,
// siftDown, siftUp, peek, remove, and insert methods.
// Feel free to add new properties and methods to the class.
type MinHeap []int

func NewMinHeap(array []int) *MinHeap {
	// Do not edit the lines below.
	heap := MinHeap(array)
	ptr := &heap
	ptr.BuildHeap(array)
	return ptr
}

func (h *MinHeap) BuildHeap(array []int) {
	copy(*h, array)
	firstPIdx := h.getParentIdx(len(array) - 1)
	for i := firstPIdx; i >= 0; i-- {
		h.siftDown(i, h.getLen())
	}
}

func (h MinHeap) Peek() int {
	return h[0]
}

func (h *MinHeap) Insert(value int) {
	(*h) = append((*h), value)
	h.siftUp()
}

func (h *MinHeap) Remove() int {
	// swap the 0th idx with the last
	h.swap(0, h.getLen())
	// pop the last child
	valueToRemove := (*h)[h.getLen()]
	(*h) = (*h)[:h.getLen()-1]
	// then call siftDown
	h.siftDown(0, h.getLen())
	return valueToRemove
}

func (h *MinHeap) siftDown(currIdx, endIndex int) {
	c1Idx, c2Idx := h.getChildIdxPair(currIdx)
	for c1Idx <= endIndex {
		if (currIdx*2)+2 > endIndex {
			c2Idx = -1
		}
		idxToSwap := -1
		if c2Idx != -1 && (*h)[c2Idx] < (*h)[c1Idx] {
			idxToSwap = c2Idx
		} else {
			idxToSwap = c1Idx
		}
		if (*h)[idxToSwap] < (*h)[currIdx] {
			h.swap(currIdx, idxToSwap)
			currIdx = idxToSwap
			c1Idx, c2Idx = h.getChildIdxPair(currIdx)
		} else {
			break
		}
	}
}

func (h *MinHeap) siftUp() {
	l_heap := len((*h)) - 1
	currIdx := l_heap
	pIdx := h.getParentIdx(currIdx)
	for currIdx > 0 && (*h)[currIdx] < (*h)[pIdx] {
		h.swap(currIdx, pIdx)
		currIdx = pIdx
		pIdx = h.getParentIdx(currIdx)
	}
}

func (h *MinHeap) swap(toSwapIdx, swapWithIdx int) {
	tmp := (*h)[swapWithIdx]
	(*h)[swapWithIdx] = (*h)[toSwapIdx]
	(*h)[toSwapIdx] = tmp
}

func (h *MinHeap) getLen() int {
	// zero index based length
	return len(*h) - 1
}

// other helpers

func (h *MinHeap) getParentIdx(childsIdx int) int {
	return int(math.Floor((float64(childsIdx-1) / 2)))
}

func (h *MinHeap) getChildIdxPair(parentIdx int) (int, int) {
	_i := 2 * parentIdx
	return _i + 1, _i + 2
}
