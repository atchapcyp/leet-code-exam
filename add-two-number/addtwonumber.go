package addtwo

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	fmt.Println("Add ", l1.Val, l2.Val)
	fmt.Printf(" l1 : %+v\n", l1)
	fmt.Printf(" l2 : %+v\n", l2)
	if l1.Next == nil && l2.Next == nil {
		if l1.Val+l2.Val > 9 {
			return &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  (l1.Val + l2.Val) % 10,
					Next: nil,
				},
			}
		}
		return &ListNode{
			Val:  l1.Val + l2.Val,
			Next: nil,
		}
	} else if l1.Next == nil && l2.Next != nil {
		return AddTwoNumbers(l1, l2.Next)
	} else if l1.Next != nil && l2.Next == nil {
		return AddTwoNumbers(l1.Next, l2)
	}
	return &ListNode{
		Val:  (l1.Val + l2.Val) % 10,
		Next: AddTwoNumbers(l1.Next, l2.Next)}
}

func ExactSameNode(l1 *ListNode, l2 *ListNode) bool {
	if l1.Next == nil && l2.Next == nil {
		return l1.Val == l2.Val
	} else if l1.Next == nil && l2.Next != nil {
		return false
	} else if l1.Next != nil && l2.Next == nil {
		return false
	}
	return ExactSameNode(l1.Next, l2.Next)
}
