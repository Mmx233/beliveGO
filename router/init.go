package router

import (
	"embed"
	"github.com/Mmx233/beliveGO/middlewares"
	"github.com/gin-gonic/gin"
)

var E *gin.Engine

func Init(FS embed.FS) error {
	gin.SetMode(gin.ReleaseMode)
	E = gin.Default()
	E.Use(middlewares.Secure(), middlewares.GZIP())
	G := E.Group("/api")
	G.Use(middlewares.Cors(), middlewares.Options())
	routerStatic(FS)
	routerUser(G.Group("/user"))

	return E.Run(":80")
}
