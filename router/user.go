package router

import (
	controllers "github.com/Mmx233/beliveGO/controllers/data"
	"github.com/gin-gonic/gin"
)

func routerUser(G *gin.RouterGroup) {
	G.GET("/avatar", controllers.Avatar)
}
