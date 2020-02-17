package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	artist "./artist"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/artists", handleArtists)
	http.ListenAndServe(":8080", nil)
}
func handleArtists(w http.ResponseWriter, req *http.Request) {
	artists := getArtists()
	a := artist.Artists{Artists: artists}
	tmpl := template.Must(template.ParseFiles("./artist/artists.html"))
	tmpl.Execute(w, a)
	// w.Header().Set("Content-Type", "text/html")

	// body, err := artists.Marshal()
	// // if err != nil {
	// // 	w.WriteHeader(400)
	// // 	fmt.Fprintf(w, "Нормально данные вводи, sumelek")
	// // }
	// w.Write(body)

}

func getArtists() []artist.Artist {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	artists, err := artist.UnmarshalArtist(body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(artists[0])
	return artists
}
