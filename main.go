package main

import (
	"embed"
	"github.com/Mmx233/beliveGO/router"
	"log"
)

//go:embed frontend/build/*
var FS embed.FS

func main() {
	log.Println("Sys init success")
	if err := router.Init(FS); err != nil {
		log.Println(err)
	}
}
