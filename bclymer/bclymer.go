package bclymer

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "bclymer/static/index.html")
}

func StartServer(urlPrefix string) {
	if urlPrefix != "" {
		urlPrefix = "/" + urlPrefix
	}
	http.HandleFunc("/", handler)

	http.Handle(urlPrefix+"/static/", http.StripPrefix(urlPrefix+"/static", http.FileServer(http.Dir("bclymer/static"))))
	log.Println("bClymer is running...")
}
