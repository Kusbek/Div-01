package artist

import (
	"fmt"
	"strings"
	"sync"
)

func Search(artists []Artist, keyword string) []string {

	var a []string
	c := make(chan string, len(artists))
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

func search(index int, c chan string, keyword string, artist Artist, wg *sync.WaitGroup) {
	defer wg.Done()
	if contains(artist.Name, keyword) {
		c <- artist.Name + " - " + "artist/band name"
		return
	}

	for _, member := range artist.Members {
		if contains(member, keyword) {
			c <- member + " - " + "member"
			return
		}
	}

	if contains(artist.FirstAlbum, keyword) {
		c <- artist.FirstAlbum + " - " + "first album date"
		return
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
