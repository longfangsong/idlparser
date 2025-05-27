package typeref

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseTypeName(t *testing.T) {
	code := "MyType"
	result := ParseTypeName(code)
	assert.Equal(t, "MyType", result.Output.Name)
}
