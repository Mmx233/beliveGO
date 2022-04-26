package main

import (
	"github.com/Mmx233/beliveGO/router"
	"log"
)

func main() {
	log.Println("Sys init success")
	if err := router.E.Run(":80"); err != nil {
		log.Println(err)
	}
}
