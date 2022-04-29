package main

import (
	"embed"
	"github.com/Mmx233/beliveGO/router"
	"github.com/gin-gonic/gin"
	"io"
	"io/fs"
	"log"
	"net/http"
	"strings"
)

//go:embed frontend/build/*
var FS embed.FS

func init() {
	fe, e := fs.Sub(FS, "frontend/build")
	if e != nil {
		log.Fatal(e)
	}
	fileServer := http.StripPrefix("/", http.FileServer(http.FS(fe)))
	router.E.Use(func(c *gin.Context) {
		path := strings.TrimPrefix(c.Request.URL.Path, "/")
		if path == "" || c.FullPath() == "" {
			path = "index.html"
		}
		f, e := fe.Open(path)
		if e != nil {
			if _, ok := e.(*fs.PathError); ok {
				return
			}
			c.AbortWithStatus(500)
			log.Println(e)
			return
		}
		_ = f.Close()
		fileServer.ServeHTTP(c.Writer, c.Request)
		c.Abort()
	})
	router.E.NoRoute(func(c *gin.Context) {
		f, e := fe.Open("index.html")
		if e != nil {
			log.Println(e)
			return
		}
		defer f.Close()
		_, _ = io.Copy(c.Writer, f)
	})
}
