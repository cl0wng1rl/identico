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
	message := getMessageParameter(w, r)
	quadrantSize := getQuadrantSizeParameter(r)

	id := identicon.QuadrantIdenticon{}
	s := id.Create(message, quadrantSize)
	w.Header().Set("Content-Type", "image/svg+xml")
	if message != "" {

	}
	io.WriteString(w, s.ToString())
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	http.HandleFunc("/", welcome)
	http.ListenAndServe(":"+port, nil)
}

func getMessageParameter(w http.ResponseWriter, r *http.Request) string {
	message := getURLParameter("message", "", r)
	if message == "" {
		w.Header().Set("Cache-Control", "no-store")
		return fmt.Sprintf("%v", time.Now().UnixNano())
	}
	return message
}

func getQuadrantSizeParameter(r *http.Request) int {
	size, err := strconv.Atoi(getURLParameter("quadrant", "8", r))
	if err != nil {
		log.Fatal(err)
	}
	if size > 128 {
		size = 128
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
