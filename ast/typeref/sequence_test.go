package typeref

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseSequence(t *testing.T) {
	code := "sequence<byte>"
	result := ParseSequence(code)
	_, isByte := result.Output.InnerType.(ByteType)
	if !isByte {
		t.Errorf("Expected ByteType, got %T", result.Output)
	}
	code = "sequence<T>"
	result = ParseSequence(code)
	inner, isTypeName := result.Output.InnerType.(TypeName)
	if !isTypeName {
		t.Errorf("Expected TypeName, got %T", result.Output)
	}
	assert.Equal(t, "T", inner.Name)
}
