// main.go

package main

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gabrielbarker/identico/identicon"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	currentTimeString := time.Now().Format("2006-01-02 15:04:05 Monday")
	id := identicon.SimpleIdenticon{}
	s := id.Create(currentTimeString)
	w.Header().Set("Content-Type", "image/svg+xml")
	io.WriteString(w, s.ToString())
}

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", welcome)
	http.ListenAndServe(":"+port, nil)
}
