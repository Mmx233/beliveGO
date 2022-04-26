package main

import (
	"embed"
	"github.com/Mmx233/beliveGO/router"
	"github.com/gin-gonic/gin"
	"io"
	"io/fs"
	"log"
)

//go:embed frontend/*
var FS embed.FS

func init() {
	router.E.GET("/*path", func(c *gin.Context) {
		path := c.Request.URL.Path
	start:
		path = "frontend/build" + path
		f, e := FS.Open(path)
		if e != nil {
			if _, ok := e.(*fs.PathError); ok {
				path = "/index.html"
				goto start
			}

			log.Println(e)
			c.Status(500)
			return
		}
		defer f.Close()

		_, _ = io.Copy(c.Writer, f)
		c.Status(200)
	})
}
