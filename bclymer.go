package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func helpTime(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, time.Now().Unix())
}

func reportIP(w http.ResponseWriter, r *http.Request) {
	file, err := os.OpenFile("rasp.txt", os.O_RDWR|os.O_APPEND, 0660)
	defer file.Close()
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	file.WriteString(string(body))
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/currentTime.php", helpTime)
	http.HandleFunc("/reportip", reportIP)

	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
	log.Println("bClymer is running...")
	http.ListenAndServe(":42125", nil)
}
