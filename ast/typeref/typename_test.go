package typeref

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseTypeName(t *testing.T) {
	code := "MyType"
	result := ParseTypeName(code)
	assert.Equal(t, "MyType", result.Output.Name)

	code = "MyType>"
	result = ParseTypeName(code)
	fmt.Println(result)
}
