// main.go

package main

import (
	"io"
	"net/http"
	"os"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>Welcome to Identico!</h1>")
}

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", welcome)
	http.ListenAndServe(":"+port, nil)
}
