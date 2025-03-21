package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func binarySearch(list []*ListNode, target int) int {
	left, right := 0, len(list)

	if left < right {
		mid := left + ((right - left) / 2)

		if list[mid].Val < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

func mergeKLists(lists []*ListNode) *ListNode {
	node := &ListNode{}

	if len(lists) == 0 {
		return &ListNode{}
	}
	if len(lists) == 1 && lists[0].Next == nil {
		return lists[0]
	}

	for _, elem := range lists {
		node.Next = &ListNode{
			Val: elem.Val,
		}
	}

	return node
}

func main() {
	mergeKLists(make([]*ListNode, 0))
}
