package main

import (
	"fmt"
	"net/http"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "はろわるど")
	data := "ハロワルド"
    fmt.Fprintf(w, data)
}

func main() {
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(":9000", nil)
}
