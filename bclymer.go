package main

import (
	"bclymer/aboutme"
	"bclymer/bclymer"
	"bclymer/quarto"
	"log"
	"net/http"
)

func main() {
	bclymer.StartServer("bclymer")

	qRedis := quarto.StartServer("quarto")
	defer qRedis.Quit()

	aRedis := aboutme.StartServer("me")
	defer aRedis.Quit()

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Print("ListenAndServe:", err)
	}
}
