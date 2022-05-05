package router

import (
	"embed"
	"github.com/gin-gonic/gin"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func routerStatic(FS embed.FS) {
	fe, e := fs.Sub(FS, "frontend/build")
	if e != nil {
		log.Fatal(e)
	}
	file, err := fe.Open("index.html")
	if e != nil {
		log.Fatal(e)
	}
	fileContentBytes, e := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(e)
	}
	_ = file.Close()
	index := string(fileContentBytes)

	fileServer := http.StripPrefix("/", http.FileServer(http.FS(fe)))
	E.Use(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/api") {
			return
		}

		f, e := fe.Open(strings.TrimPrefix(c.Request.URL.Path, "/"))
		if e != nil {
			if _, ok := e.(*fs.PathError); ok {
				c.Header("Content-Type", "text/html")
				c.String(200, index)
				c.Abort()
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

	E.NoRoute(func(c *gin.Context) {
		f, e := fe.Open("index.html")
		if e != nil {
			log.Println(e)
			return
		}
		defer f.Close()
		_, _ = io.Copy(c.Writer, f)
	})
}
