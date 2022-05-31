package middlewares

import (
	"github.com/Mmx233/beliveGO/controllers"
	"github.com/Mmx233/secure"
	"github.com/gin-gonic/gin"
)

func Secure() gin.HandlerFunc {
	return secure.New(&secure.Config{
		CallBack: func(c *gin.Context) {
			controllers.CallBack.Error(c, 3, nil)
		},
		MinLimitMode: true,
	})
}

func ApiSecure() gin.HandlerFunc {
	return secure.New(&secure.Config{
		CallBack: func(c *gin.Context) {
			controllers.CallBack.Error(c, 3, nil)
		},
		RateLimit:    60,
		MinLimitMode: true,
	})
}
