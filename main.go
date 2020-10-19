// main.go

package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gabrielbarker/identico/svg"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	s := svg.New()
	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			s.AddElement(svg.NewRandomSquare(i, j))
		}
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	io.WriteString(w, s.ToString())
}

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", welcome)
	http.ListenAndServe(":"+port, nil)
}
