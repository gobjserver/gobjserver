package entity_test

import (
	"testing"

	"github.com/gobjserver/gobjserver/core/entity"
	"github.com/stretchr/testify/assert"
)

// TestCreateEntity .
func TestCreateEntity(test *testing.T) {
	assert.NotNil(test, entity.Object{})
}
