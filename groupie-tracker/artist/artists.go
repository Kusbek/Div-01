package artist

import "encoding/json"

type Artists struct {
	Artists []Artist
}

func UnmarshalArtist(data []byte) ([]Artist, error) {
	var r []Artist
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
