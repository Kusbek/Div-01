package artist

import (
	"net/http"
	"strconv"
	"sync"
)

func getCreatedDateFilterDates(req *http.Request) (int64, int64) {
	dateFrom, _ := strconv.Atoi(req.URL.Query().Get("cr_date_from"))
	dateTo, _ := strconv.Atoi(req.URL.Query().Get("cr_date_to"))
	if dateTo == 0 {
		dateTo = 2100
	}
	return int64(dateFrom), int64(dateTo)
}

func getFirstAlbumFilterDates(req *http.Request) (int64, int64) {
	dateFrom, _ := strconv.Atoi(req.URL.Query().Get("fa_date_from"))
	dateTo, _ := strconv.Atoi(req.URL.Query().Get("fa_date_to"))
	if dateTo == 0 {
		dateTo = 2100
	}
	return int64(dateFrom), int64(dateTo)
}
func getCountryFromURL(req *http.Request) string {
	country := req.URL.Query().Get("country")
	if country == "" {
		return "all"
	}
	return country
}

func single(req *http.Request) bool {
	single := req.URL.Query().Get("single")

	if single == "false" {
		return false
	}
	return true
}

func getCountries(id string) []string {
	locations := getLocations(id).Locations
	var countries []string
	for _, location := range locations {
		countries = append(countries, split(location, "-")[1])
	}
	return countries
}

func countryInCountries(country string, countries []string) bool {
	for _, c := range countries {
		if c == country {
			return true
		}
	}
	return false
}

func filter(index int, c chan Artist, req *http.Request, artist Artist, wg *sync.WaitGroup) {
	defer wg.Done()
	intoChannel := true
	//Filter by date
	dateFrom, dateTo := getCreatedDateFilterDates(req)
	if artist.CreationDate < dateFrom || artist.CreationDate > dateTo {
		intoChannel = false

	}
	//Filter by album date
	dateFrom, dateTo = getFirstAlbumFilterDates(req)
	date := parseDate(artist.FirstAlbum)
	if date < dateFrom || date > dateTo {
		intoChannel = false
	}
	//Filter by checkbox filter
	if single(req) && len(artist.Members) != 1 {
		intoChannel = false
	}
	country := getCountryFromURL(req)
	countries := getCountries(strconv.Itoa(int(artist.ID)))
	if !countryInCountries(country, countries) {
		intoChannel = false
	}
	if intoChannel {
		// fmt.Println("inc", artist.Name)
		c <- artist
	}
}

func Filter(artists []Artist, req *http.Request) []Artist {
	var a []Artist
	c := make(chan Artist, len(artists))
	wg := sync.WaitGroup{}

	for i, artist := range artists {
		wg.Add(1)
		go filter(i, c, req, artist, &wg)
	}

	wg.Wait()
	close(c)

	for i := range c {
		a = append(a, i)
	}
	return a
}
