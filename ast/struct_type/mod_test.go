package struct_type

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsing(t *testing.T) {
	code := `@ann(a = b, c = d)`
	result := parseAnnotation(code)
	assert.Equal(t, "ann", result.Output.Name)
	assert.Equal(t, "b", result.Output.Values["a"])
	assert.Equal(t, "d", result.Output.Values["c"])
}
