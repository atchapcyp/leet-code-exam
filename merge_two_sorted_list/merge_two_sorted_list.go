package merge_two_sorted_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var resultNode *ListNode
	for l1 != nil || l2 != nil {
		if l1 == nil {
			resultNode = resultNode.AddNode(l2.Val)
			l2 = l2.Next
			continue
		}

		if l2 == nil {
			resultNode = resultNode.AddNode(l1.Val)
			l1 = l1.Next
			continue
		}

		if l1.Val <= l2.Val {
			resultNode = resultNode.AddNode(l1.Val)
			l1 = l1.Next
		} else {
			resultNode = resultNode.AddNode(l2.Val)
			l2 = l2.Next
		}
	}
	return resultNode
}

func (ln *ListNode) AddNode(val int) *ListNode {
	n := &ListNode{Val: val, Next: nil}

	if ln == nil {
		ln = n
		return ln
	}

	if ln.Next == nil {
		ln.Next = n
	} else {
		cur := ln
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = n
	}
	return ln
}
