package artist

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Locations struct {
	ID        int64    `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

func UnmarshalLocations(data []byte) (Locations, error) {
	var r Locations
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Locations) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func getLocations(id string) Locations {
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/locations/%s", id)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	relations, err := UnmarshalLocations(body)
	if err != nil {
		fmt.Println(err)
	}
	return relations
}
