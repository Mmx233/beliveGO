package main

import (
	"github.com/Mmx233/beliveGO/router"
	"log"
)

func main() {
	if err := router.E.Run(":80"); err != nil {
		log.Println(err)
	}
}
