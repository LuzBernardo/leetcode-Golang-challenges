package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	filtered := make([]*ListNode, 0, len(lists))

	for _, node := range lists {
		if node != nil {
			filtered = append(filtered, node)
		}
	}

	if len(filtered) == 0 {
		return nil
	}
	if len(filtered) == 1 {
		return filtered[0]
	}

	h := NewMinHeap()

	for _, node := range filtered {
		h.Push(node)
	}

	dummy := &ListNode{}
	curr := dummy
	for h.Len() > 0 {
		node := h.Pop()
		curr.Next = node
		curr = curr.Next
		if node.Next != nil {
			h.Push(node.Next)
		}
	}

	return dummy.Next
}

func main() {
	lists := []*ListNode{
		CreateLinkedList([]int{1, 4, 5}),
		CreateLinkedList([]int{1, 3, 4}),
		CreateLinkedList([]int{2, 6}),
	}
	result := mergeKLists(lists)
	PrintList(result)
}
