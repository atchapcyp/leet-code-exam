package validparent

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsValidparent(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"{}[]", true},
		{"{[]}", true},
		{"", true},
		{"{{}}{]", false},
		{"{", false},
		{"{{", false},
		{"]", false},
		{"]]", false},
	}

	for _, test := range tests {
		require.Equal(t, test.expected, IsValid(test.input), test.input)
	}
}
