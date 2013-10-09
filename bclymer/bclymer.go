package bclymer

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "bclymer/static/index.html")
}

func helpTime(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, time.Now().Unix())
}

func StartServer(urlPrefix string) {
	if urlPrefix != "" {
		urlPrefix = "/" + urlPrefix
	}
	http.HandleFunc("/", handler)
	http.HandleFunc("/currentTime.php", helpTime)

	http.Handle(urlPrefix+"/static/", http.StripPrefix(urlPrefix+"/static", http.FileServer(http.Dir("bclymer/static"))))
	log.Println("bClymer is running...")
}
