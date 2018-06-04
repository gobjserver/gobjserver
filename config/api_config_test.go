package config_test

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gobjserver/gobjserver/config"
)

// SetupRouter .
func SetupRouter() *gin.Engine {
	router := gin.Default()
	gin.SetMode(gin.TestMode)
	return router
}

// TestSetAPIPaths .
func TestSetAPIPaths(test *testing.T) {
	router := SetupRouter()

	config.SetAPIPaths(router)

}
