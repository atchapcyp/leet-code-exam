package longestcmn

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReverseInt(t *testing.T) {
	tests := []struct {
		input  []string
		result string
	}{
		{[]string{"flower", "flow", "flur"}, "fl"},
		{[]string{"flower"}, "flower"},
		{[]string{"fl"}, "fl"},
		{[]string{""}, ""},
	}
	for _, test := range tests {
		actual := LongestCommonPrefix(test.input)
		require.Equal(t, test.result, actual)
	}
}
