package main

import (
	"embed"
	"github.com/Mmx233/beliveGO/router"
	"github.com/gin-gonic/gin"
	"io"
	"io/fs"
	"log"
	"net/http"
)

//go:embed frontend/build/*
var FS embed.FS

func init() {
	fe, e := fs.Sub(FS, "frontend/build")
	if e != nil {
		log.Fatal(e)
	}
	router.E.StaticFS("/", http.FS(fe))
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
