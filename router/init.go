package router

import (
	"github.com/Mmx233/beliveGO/middlewares"
	"github.com/gin-gonic/gin"
)

var E *gin.Engine

func init() {
	E = gin.Default()
	E.Use(middlewares.Secure())
	G := E.Group("/api")

	E.GET("/")
}
