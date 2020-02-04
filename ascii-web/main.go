package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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

	result := getArt(text, format)
	fmt.Fprintf(w, "%v\n", result)
	// Do something
}

func getArt(text, format string) string {
	template, e := openFileAndReadLineByLine(format)
	if e {
		return "Sorry, I could not print what you have typed"
	}
	result := asciiart(text, template)
	return result
}
func asciiart(str string, lines []string) string {
	//Creating the art itself
	var arr []rune
	Newline := false
	result := ""
	for i, r := range str {
		if Newline {
			Newline = false
			result = stringArt(arr, lines, result)
			arr = []rune{}
			continue
		}

		if r == 92 && len(str) != i+1 {
			if str[i+1] == 110 {
				Newline = true
				continue
			}
		}
		arr = append(arr, r)
	}
	return stringArt(arr, lines, result)
}

func openFileAndReadLineByLine(o string) ([]string, bool) {
	file := fmt.Sprintf("%s.txt", o)
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return nil, true
	}
	lines := strings.Split(string(bytes), "\n")

	return lines, false
}

func hasError(str string) bool {
	for _, r := range str {
		if r < 32 || r > 126 {
			return true
		}
	}
	return false
}

//Make a string-art of given rune array, based on lines art
func stringArt(arr []rune, lines []string, result string) string {

	for line := 0; line < 8; line++ {
		for _, r := range arr {
			skip := (r-32)*9 + 1
			result = result + lines[line+int(skip)]
		}
		result = result + "\n"
	}
	return result
}
