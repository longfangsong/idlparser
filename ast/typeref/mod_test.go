package typeref

import (
	"testing"
)

func TestParse(t *testing.T) {
	code := "byte"
	result := ParseTypeRef(code)
	_, ok := result.Output.(ByteType)
	if !ok {
		t.Errorf("Expected ByteType, got %T", result.Output)
	}
	code = "StructTypeName"
	result = ParseTypeRef(code)
	_, ok = result.Output.(TypeName)
	if !ok {
		t.Errorf("Expected StructType, got %T", result.Output)
	}

	result = ParseTypeRef("unsigned long")
	if _, ok := result.Output.(UnsignedLongType); !ok {
		t.Errorf("Expected UnsignedLongType, got %T", result.Output)
	}

	result = ParseTypeRef("unsigned long    long")
	if _, ok := result.Output.(UnsignedLongLongType); !ok {
		t.Errorf("Expected UnsignedLongLongType, got %T", result.Output)
	}
}
