package config_test

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gobjserver/gobjserver/config"
	"github.com/stretchr/testify/assert"
)

// TestSetMiddlewares .
func TestSetMiddlewares(test *testing.T) {
	result := config.SetMiddleWares(gin.New())
	assert.NotNil(test, result)
}
