// main.go

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gabrielbarker/identico/identicon"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	message := getMessageParameter(r)
	quadrantSize := getQuadrantSizeParameter(r)

	id := identicon.SimpleIdenticon{}
	s := id.Create(message, quadrantSize)
	w.Header().Set("Content-Type", "image/svg+xml")
	w.Header().Set("Cache-Control", "no-store")
	io.WriteString(w, s.ToString())
}

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", welcome)
	http.ListenAndServe(":"+port, nil)
}

func getMessageParameter(r *http.Request) string {
	return getURLParameter("message", fmt.Sprintf("%v", time.Now().UnixNano()), r)
}

func getQuadrantSizeParameter(r *http.Request) int {
	size, err := strconv.Atoi(getURLParameter("quadrant", "8", r))
	if err != nil {
		log.Fatal(err)
	}
	return size
}

func getURLParameter(parameter string, defaultValue string, r *http.Request) string {
	params, ok := r.URL.Query()[parameter]
	if !ok || len(params[0]) < 1 {
		return defaultValue
	}
	return params[0]
}
