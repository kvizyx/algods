package binheap

import (
	"slices"

	"github.com/kvizyx/algods"
)

type (
	BinMaxHeap[K comparable] []*Node[K]

	Node[K comparable] struct {
		Key   K
		Value int64
	}
)

func NewBinMaxHeap[K comparable]() BinMaxHeap[K] {
	return make(BinMaxHeap[K], 0)
}

func (h *BinMaxHeap[K]) heapify() {
	for i := len(*h)/2 - 1; i >= 0; i-- {
		h.heapifyFrom(i)
	}
}

func (h *BinMaxHeap[K]) heapifyFrom(index int) {
	if h == nil || len(*h) == 0 {
		return
	}

	size := len(*h)

	largest := index
	left := 2*index + 1
	right := 2*index + 2

	if left < size && (*h)[left].Value > (*h)[largest].Value {
		largest = left
	}

	if right < size && (*h)[right].Value > (*h)[largest].Value {
		largest = right
	}

	if largest != index {
		algods.Swap(*h, index, largest)
		h.heapifyFrom(largest)
	}
}

func (h *BinMaxHeap[K]) Push(key K, value int64) bool {
	if h == nil {
		return false
	}

	if len(*h) == 0 {
		*h = append(*h, &Node[K]{
			Key:   key,
			Value: value,
		})
		return true
	}

	*h = append(*h, &Node[K]{
		Key:   key,
		Value: value,
	})

	h.heapify()

	return true
}

func (h *BinMaxHeap[K]) Pop(key K) bool {
	if h == nil {
		return false
	}

	for i, v := range *h {
		if v.Key == key {
			if 2*i+1 > len(*h) && 2*i+2 > len(*h) {
				*h = slices.Delete(*h, i, i+1)
				return true
			}

			algods.Swap(*h, i, len(*h)/2-1)
			*h = slices.Delete(*h, len(*h)/2-1, len(*h)/2)

			h.heapify()

			return true
		}
	}

	return false
}

func (h *BinMaxHeap[K]) Inc(key K) bool {
	if h == nil {
		return false
	}

	for i, v := range *h {
		if v.Key == key {
			(*h)[i].Value++
			h.heapify()

			return true
		}
	}

	return false
}

func (h *BinMaxHeap[K]) Desc(key K) bool {
	if h == nil {
		return false
	}

	for i, v := range *h {
		if v.Key == key {
			(*h)[i].Value--
			h.heapify()

			return true
		}
	}

	return false
}

func (h *BinMaxHeap[K]) Peek() *Node[K] {
	if h == nil || len(*h) == 0 {
		return nil
	}

	return (*h)[0]
}
