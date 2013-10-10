package bclymer

//package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	urlPrefix    = ""
	folderPrefix = "bclymer/"
)

func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, folderPrefix+"static/index.html")
}

func helpTime(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, time.Now().Unix())
}

func main() {
	http.HandleFunc(urlPrefix+"/", handler)
	http.HandleFunc(urlPrefix+"/currentTime.php", helpTime)

	http.Handle("/"+folderPrefix+"static/", http.StripPrefix("/"+folderPrefix+"static", http.FileServer(http.Dir(folderPrefix+"static"))))
	log.Println("bClymer is running...")
}

func StartServer() {
	main()
}
