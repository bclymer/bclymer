package main

import (
	"bclymer/bclymer"
	"bclymer/quarto"
	"log"
	"net/http"
)

func main() {
	quarto.StartServer("quarto")
	bclymer.StartServer("bclymer")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Print("ListenAndServe:", err)
	}
}
