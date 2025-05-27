package typeref

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseBitField(t *testing.T) {
	tests := []struct {
		input    string
		expected TypeRef
	}{
		{"bitfield<8>", BitFieldType{Width: 8}},
		{"bitfield<16>", BitFieldType{Width: 16}},
		{"bitfield<32>", BitFieldType{Width: 32}},
	}

	for _, test := range tests {
		result := ParseBitField(test.input)
		assert.Equal(t, test.expected, result.Output)
	}
}
