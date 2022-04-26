package controllers

import (
	"github.com/gin-gonic/gin"
)

type call struct{}

var CallBack call

func (a *call) Error(c *gin.Context, code uint, e error) {
	c.AsciiJSON(400, gin.H{
		"code": code,
	})
	c.Abort()
}

func (*call) Success(c *gin.Context, data interface{}) {
	c.AsciiJSON(200, gin.H{
		"code": 0,
		"data": data,
	})
}

func (a *call) Default(c *gin.Context) {
	a.Success(c, nil)
}
