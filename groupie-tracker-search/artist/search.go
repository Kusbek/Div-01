package artist

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

// func Search(artists []Artist, keyword string) []string {

// 	var a []string
// 	c := make(chan string, len(artists))
// 	wg := sync.WaitGroup{}

// 	for i, artist := range artists {
// 		wg.Add(1)
// 		go search(i, c, keyword, artist, &wg)
// 	}

// 	wg.Wait()
// 	close(c)

// 	for i := range c {
// 		a = append(a, i)
// 	}
// 	return a
// }
// func search(index int, c chan string, keyword string, artist Artist, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	if contains(artist.Name, keyword) {
// 		c <- artist.Name + " - " + "artist/band name"
// 		return
// 	}

// 	for _, member := range artist.Members {
// 		if contains(member, keyword) {
// 			c <- member + " - " + "member"
// 			return
// 		}
// 	}

// 	if contains(artist.FirstAlbum, keyword) {
// 		c <- artist.FirstAlbum + " - " + artist.Name + " first album date"
// 		return
// 	}
// 	n := strconv.Itoa(int(artist.CreationDate))
// 	if contains(n, keyword) {
// 		c <- n + " - " + artist.Name + " creation date"
// 		return
// 	}

// }
func Search(artists []Artist, keyword string) []Artist {

	var a []Artist
	c := make(chan Artist, len(artists))
	wg := sync.WaitGroup{}

	for i, artist := range artists {
		wg.Add(1)
		go search(i, c, keyword, artist, &wg)
	}

	wg.Wait()
	close(c)

	for i := range c {
		a = append(a, i)
	}
	fmt.Println(a)
	return a
}
func search(index int, c chan Artist, keyword string, artist Artist, wg *sync.WaitGroup) {
	defer wg.Done()
	if contains(artist.Name, keyword) {
		c <- artist
		return
	}

	for _, member := range artist.Members {
		if contains(member, keyword) {
			c <- artist
			return
		}
	}

	if contains(artist.FirstAlbum, keyword) {
		c <- artist
		return
	}
	n := strconv.Itoa(int(artist.CreationDate))
	if contains(n, keyword) {
		c <- artist
		return
	}

	locations := getLocations(strconv.Itoa(int(artist.ID)))

	for _, location := range locations.Locations {
		if contains(location, keyword) {
			c <- artist
			return
		}
	}

}

func contains(str, keyword string) bool {
	str = strings.ToLower(str)
	keyword = strings.ToLower(keyword)
	length := len(str) - len(keyword)
	for i := 0; i <= length; i++ {
		if str[i:i+len(keyword)] == keyword {
			return true
		}
	}
	return false
}
