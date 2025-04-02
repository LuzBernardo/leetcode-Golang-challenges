package main

type MinHeap struct {
	nodes []*ListNode
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		nodes: make([]*ListNode, 0),
	}
}

func (h *MinHeap) Len() int {
	return len(h.nodes)
}

func (h *MinHeap) heapifyUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if h.nodes[parent].Val <= h.nodes[index].Val {
			break
		}
		h.nodes[parent], h.nodes[index] = h.nodes[index], h.nodes[parent]
		index = parent
	}
}

func (h *MinHeap) heapifyDown(index int) {
	size := h.Len()

	for {
		smallest := index
		left := 2*index + 1
		right := 2*index + 2

		if left < size && (h.nodes[left].Val < h.nodes[smallest].Val) {
			smallest = left
		}
		if right < size && (h.nodes[right].Val < h.nodes[smallest].Val) {
			smallest = right
		}
		if smallest == index {
			break
		}
		h.nodes[index], h.nodes[smallest] = h.nodes[smallest], h.nodes[index]
		index = smallest
	}
}

func (h *MinHeap) Push(node *ListNode) {
	h.nodes = append(h.nodes, node)
	h.heapifyUp(len(h.nodes) - 1)
}

func (h *MinHeap) Pop() *ListNode {
	if h.Len() == 0 {
		return nil
	}

	min := h.nodes[0]
	last := h.Len() - 1

	h.nodes[0] = h.nodes[last]
	h.nodes = h.nodes[:last]

	if h.Len() > 0 {
		h.heapifyDown(0)
	}

	return min
}
