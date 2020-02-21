package artist

import (
	"fmt"
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

func filter(index int, c chan Artist, req *http.Request, artist Artist, wg *sync.WaitGroup) {
	defer wg.Done()
	intoChannel := true
	dateFrom, dateTo := getCreatedDateFilterDates(req)
	if artist.CreationDate < dateFrom || artist.CreationDate > dateTo {
		intoChannel = false

	}

	dateFrom, dateTo = getFirstAlbumFilterDates(req)
	date := parseDate(artist.FirstAlbum)
	if date < dateFrom || date > dateTo {
		intoChannel = false
	}

	if intoChannel {
		fmt.Println("inc", artist.Name)
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
