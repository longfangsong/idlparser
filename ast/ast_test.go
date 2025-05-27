package ast

import (
	"testing"

	"github.com/longfangsong/idl-parser/ast/bitset"
	"github.com/stretchr/testify/assert"
)

func TestParsing(t *testing.T) {
	code := `module spi {
    bitset IdBits {
        bitfield<4> bid; // 4 bits for bus_id
        bitfield<12> cid;  // 12 bits for can_id
    };
}`
	result := Parse(code)
	if result.Err != nil {
		t.Errorf("Parsing failed: %v", result.Err)
	}
	assert.Equal(t, result.Output.Name, "spi")
	assert.Equal(t, len(result.Output.Content), 1)
	first, ok := result.Output.Content[0].(bitset.BitSet)
	if !ok {
		t.Errorf("Expected BitSet, got %T", first)
	}
	assert.Equal(t, first.Name, "IdBits")
	assert.Equal(t, len(first.Fields), 2)
	assert.Equal(t, first.Fields[0].Name, "bid")
	assert.Equal(t, first.Fields[0].Type.Width, uint8(4))
	assert.Equal(t, first.Fields[1].Name, "cid")
	assert.Equal(t, first.Fields[1].Type.Width, uint8(12))
}
