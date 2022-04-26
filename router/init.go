package router

import (
	"github.com/Mmx233/beliveGO/middlewares"
	"github.com/gin-gonic/gin"
)

var E *gin.Engine

func init() {
	gin.SetMode(gin.ReleaseMode)
	E = gin.Default()
	E.Use(middlewares.Secure(), middlewares.GZIP())
	//G := E.Group("/api")
}
