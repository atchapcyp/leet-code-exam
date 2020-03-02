package twosum

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTwosum(t *testing.T) {

	tests := []struct {
		input  []int
		target int
		result []int
	}{
		{[]int{1, 2, 3, 4}, 3, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{4, 5, 1, 3, 5, 2, 1}, 2, []int{2, 6}},
	}

	for _, test := range tests {
		actual := TwoSum(test.input, test.target)
		require.Equal(t, test.result, actual)
	}
}
