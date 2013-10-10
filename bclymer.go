package main

import (
	"bclymer/aboutme"
	"bclymer/bclymer"
	"bclymer/quarto"
	"log"
	"net/http"
)

func main() {
	bclymer.StartServer()

	qRedis := quarto.StartServer("quarto")
	defer qRedis.Quit()

	aboutme.StartServer()

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Print("ListenAndServe:", err)
	}
}
