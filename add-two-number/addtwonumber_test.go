package addtwo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddTwoNumber(t *testing.T) {

	tests := []struct {
		input_1 *ListNode
		input_2 *ListNode
		result  *ListNode
	}{
		{nodeFromNum(123), nodeFromNum(456), nodeFromNum(579)},
		{nodeFromNum(342), nodeFromNum(465), nodeFromNum(807)},
		{nodeFromNum(348), nodeFromNum(465), nodeFromNum(813)},
		{nodeFromNum(123), nodeFromNum(7), nodeFromNum(130)},
		{nodeFromNum(243), nodeFromNum(564), nodeFromNum(807)},
		{nodeFromNum(99), nodeFromNum(1), nodeFromNum(100)},
		{nodeFromNum(18), nodeFromNum(0), nodeFromNum(18)},
		{nodeFromNum(0), nodeFromNum(0), nodeFromNum(0)},
	}

	for _, test := range tests {
		actual := AddTwoNumbers(test.input_1, test.input_2)
		fmt.Printf(" test 1: %+v\n", test.input_1)
		fmt.Printf(" test 2: %+v\n", test.input_2)
		fmt.Printf(" result: %+v\n", test.result)
		fmt.Printf(" actual1: %+v\n", actual)
		fmt.Printf(" actual2: %+v\n", actual.Next)
		fmt.Printf(" actual3: %+v\n", actual.Next.Next)
		require.True(t, ExactSameNode(actual, test.result))

	}
}

func nodeFromNum(n int) *ListNode {
	firstDigit := n % 10
	secondDigit := ((n - firstDigit) % 100) / 10
	thirdDigit := ((n - secondDigit - firstDigit) % 1000) / 100

	n1 := &ListNode{
		Val:  thirdDigit,
		Next: nil,
	}

	n2 := &ListNode{
		Val:  secondDigit,
		Next: n1,
	}

	n3 := &ListNode{
		Val:  firstDigit,
		Next: n2,
	}

	return n3
}
