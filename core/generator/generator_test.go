package generator_test

import (
	"testing"

	"github.com/gobjserver/gobjserver/core/generator"
	"github.com/stretchr/testify/assert"
)

// TestCurrentDateTime.
func TestCurrentDateTime(test *testing.T) {
	assert.NotNil(test, generator.CurrentDateTime())
}
