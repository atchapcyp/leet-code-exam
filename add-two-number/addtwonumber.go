package addtwo

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry bool
	carryNode := &ListNode{
		Val:  1,
		Next: nil,
	}
	var l3 *ListNode
	if l2.Val+l1.Val > 9 {
		carry = true
	}

	if l1.Next == nil && l2.Next == nil {
		if carry {
			l3 = &ListNode{
				Val:  (l1.Val + l2.Val) % 10,
				Next: carryNode,
			}
		} else {
			l3 = &ListNode{
				Val:  l1.Val + l2.Val,
				Next: nil,
			}
		}
		return l3
	} else if l1.Next == nil && l2.Next != nil {
		if carry {
			l3 = AddTwoNumbers(l2.Next, carryNode)
		} else {
			l3 = l2.Next
		}
	} else if l1.Next != nil && l2.Next == nil {
		if carry {
			l3 = AddTwoNumbers(l1.Next, carryNode)
		} else {
			l3 = l1.Next
		}
	} else {
		l3 = AddTwoNumbers(l1.Next, l2.Next)
		if carry {
			l3 = AddTwoNumbers(l3, carryNode)
		}
	}

	return &ListNode{
		Val:  (l1.Val + l2.Val) % 10,
		Next: l3,
	}
}

func ExactSameNode(l1 *ListNode, l2 *ListNode) bool {
	fmt.Printf("Compare %d , %d \n", l1.Val, l2.Val)
	if l1.Val != l2.Val {
		return false
	}
	if l1.Next == nil && l2.Next == nil {
		return l1.Val == l2.Val
	} else if l1.Next == nil && l2.Next != nil {
		return false
	} else if l1.Next != nil && l2.Next == nil {
		return false
	} else {
		return ExactSameNode(l1.Next, l2.Next)
	}
}
