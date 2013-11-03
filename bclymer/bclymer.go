package bclymer

//package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
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

func reportIP(w http.ResponseWriter, r *http.Request) {
	file, err := os.OpenFile("rasp.txt", O_RDWR|O_APPEND, 0660)
	defer file.Close()
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	file.WriteString(body + "\n")
}

func main() {
	http.HandleFunc(urlPrefix+"/", handler)
	http.HandleFunc(urlPrefix+"/currentTime.php", helpTime)
	http.HandleFunc(urlPrefix+"/reportip", reportIP)

	http.Handle("/"+folderPrefix+"static/", http.StripPrefix("/"+folderPrefix+"static", http.FileServer(http.Dir(folderPrefix+"static"))))
	log.Println("bClymer is running...")
}

func StartServer() {
	main()
}
