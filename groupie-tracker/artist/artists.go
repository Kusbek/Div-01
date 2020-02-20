package artist

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type Artists struct {
	Artists []Artist
}

func UnmarshalArtist(data []byte) (Artist, error) {
	var r Artist
	err := json.Unmarshal(data, &r)
	return r, err
}
func UnmarshalArtists(data []byte) ([]Artist, error) {
	var r []Artist
	err := json.Unmarshal(data, &r)
	return r, err
}

func UnmarshalRelations(data []byte) (Relations, error) {
	var r Relations
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Artists) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Artist struct {
	ID           int64    `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int64    `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type ArtistTemplate struct {
	Artist    Artist
	Relations Relations
}

type Relations struct {
	ID             int64               `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

func HandleArtist(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	artist := getArtist(id)
	relations := getRelations(id)
	var intoTemplate ArtistTemplate = ArtistTemplate{Artist: artist, Relations: relations}
	tmpl := template.Must(template.ParseFiles("./artist/artist.html"))
	tmpl.Execute(w, intoTemplate)
}
func getRelations(id string) Relations {
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/relation/%s", id)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	relations, err := UnmarshalRelations(body)
	if err != nil {
		fmt.Println(err)
	}
	return relations
}
func getArtist(id string) Artist {
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%s", id)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	artist, err := UnmarshalArtist(body)
	if err != nil {
		fmt.Println(err)
	}
	return artist
}

func HandleArtists(w http.ResponseWriter, req *http.Request) {
	artists := getArtists()
	a := Artists{Artists: artists}
	tmpl := template.Must(template.ParseFiles("./artist/artists.html"))
	tmpl.Execute(w, a)
}

func getArtists() []Artist {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	artists, err := UnmarshalArtists(body)
	if err != nil {
		fmt.Println(err)
	}
	return artists
}
