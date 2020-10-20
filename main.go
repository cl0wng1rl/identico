// main.go

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gabrielbarker/identico/identicon"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	currentTimeString := fmt.Sprintf("%v", time.Now().UnixNano())
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
