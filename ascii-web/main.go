package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/asciiart", createASCII)
	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.ListenAndServe(":8080", nil)
}

func createASCII(w http.ResponseWriter, req *http.Request) {
	queries := req.URL.Query()
	text := queries.Get("text")
	format := queries.Get("format")
	if text == "" || format == "" {
		fmt.Fprintf(w, "Нормально данные вводи, ушлепок")
		return
	}
	fmt.Fprintf(w, "%v, %v\n", text, format)
	// Do something
}
