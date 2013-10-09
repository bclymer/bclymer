package main

import (
	"bclymer/aboutme"
	"bclymer/bclymer"
	"bclymer/quarto"
	"log"
	"net/http"
)

func main() {
	quarto.StartServer("quarto")
	bclymer.StartServer("bclymer")
	aboutme.StartServer("me")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Print("ListenAndServe:", err)
	}
}
