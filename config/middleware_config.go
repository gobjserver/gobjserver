package config

import (
	"github.com/gin-gonic/gin"
	"github.com/gobjserver/gobjserver/middleware"
)

// SetMiddleWares setups middlewares.
func SetMiddleWares(router *gin.Engine) gin.IRoutes {
	iroutes := router.Use(middleware.Cors())
	iroutes = router.Use(middleware.Gzip())
	iroutes = router.Use(middleware.Static())
	return iroutes
}
