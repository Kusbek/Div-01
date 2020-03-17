package main

import (
	"DIV-01/groupie-tracker-search/artist"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/artists", artist.HandleArtists)
	http.HandleFunc("/artist", artist.HandleArtist)
	http.ListenAndServe(":8080", nil)
}
