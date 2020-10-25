package merge_two_sorted_list

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMergeTwoSortedList(t *testing.T) {
	tests := []struct {
		input1 []int
		input2 []int
		result []int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, []int{7, 8, 9, 10}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{[]int{7, 8, 9, 10}, []int{1, 2, 3, 4, 5, 6}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{[]int{1, 2, 3, 4, 5, 7}, []int{6, 8, 9, 10}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{[]int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
		{[]int{0}, []int{0}, []int{0, 0}},
	}
	for _, test := range tests {
		n1 := buildListNode(test.input1)
		n2 := buildListNode(test.input2)
		expected := buildListNode(test.result)
		actual := MergeTwoLists(n1, n2)
		require.Equal(t, expected, actual)
		require.True(t, compareListNode(actual, expected))
	}
}

func buildListNode(l []int) *ListNode {
	ln0 := &ListNode{l[0], nil}
	for i := 1; i < len(l); i++ {
		ln0.AddNode(l[i])
	}
	return ln0
}

func compareListNode(x *ListNode, y *ListNode) bool {
	fmt.Printf("X : %d , Y : %d\n", x.Val, y.Val)
	if x.Val != y.Val {
		return false
	}
	if x.Next == nil || y.Next == nil {
		return x.Next == y.Next
	}
	return compareListNode(x.Next, y.Next)
}
