package heap

import (
	"time"

	"github.com/algorithmsRepeat/soft"
)

type Heap struct {
	size int
	*soft.Soft
}

func NewHeap(arr []int) *Heap {
	h := &Heap{
		size: 0,
		Soft: soft.NewSoft(arr),
	}
	h.size = len(arr)
	h.heapify()
	return h
}

func (h *Heap) SortTest() {
	before := time.Now()
	h.Sort()
	end := time.Now()
	timeStr := end.Sub(before).String()
	h.String("heap sort", timeStr)
}

func (h *Heap) siftDown(index int) {
	node := h.Array[index]
	for {
		maxIndex := h.findMax(index)
		if maxIndex < 0 || node >= h.Array[maxIndex] {
			break
		}
		h.Array[index] = h.Array[maxIndex]
		index = maxIndex
	}
	h.Array[index] = node
}

func (h *Heap) findMax(index int) int {
	left := index<<1 + 1
	if left >= h.size {
		return -1
	}
	if left == h.size-1 {
		return left
	} else {
		if h.Array[left] > h.Array[left+1] {
			return left
		}
		return left + 1
	}
}

func (h *Heap) heapify() {
	for index := h.size >> 1; index >= 0; index-- {
		h.siftDown(index)
	}

}

func (h *Heap) Sort() {
	for h.size > 1 {
		h.size--
		h.Swap(0, h.size)
		h.siftDown(0)
	}
}
