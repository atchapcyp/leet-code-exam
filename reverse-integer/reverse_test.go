package reverseint

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReverseInt(t *testing.T) {
	tests := []struct {
		input  int
		result int
	}{
		{12, 21},
		{100, 1},
		{120, 21},
		{-12, -21},
		{-120, -21},
		{1534236469, 0},
		{-2147483648, 0},
	}
	for _, test := range tests {
		actual := Reverse(test.input)
		require.Equal(t, test.result, actual)
	}
}
