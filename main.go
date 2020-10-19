// main.go

package server

import (
	"io"
	"net/http"
	"os"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Welcome to Identico!")
}

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", welcome)
	http.ListenAndServe(":"+port, nil)
}
