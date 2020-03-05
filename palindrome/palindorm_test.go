package palindrome

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReverseInt(t *testing.T) {
	tests := []struct {
		input  int
		result bool
	}{
		{123321, true},
		{1221, true},
		{1111, true},
		{22, true},
		{123456654321, true},
		{1, true},
		{121, true},
		{-121, false},
		{-121123123, false},
		{548123487817, false},
	}
	for _, test := range tests {
		actual := IsPalindrome2(test.input)
		require.Equal(t, test.result, actual)
	}
}
