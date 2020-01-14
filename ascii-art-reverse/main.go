package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	template, e := openFileAndReadLineByLine("standard")
	if e {
		return
	}
	segment(template)
}

func segment(template []string) /*[][]string*/ {
	// result := make([][]string, 126-32)

	for i := 0; i < len(template); i = i + 9 {
		temp := template[i : i+9]
		for _, v := range temp {
			fmt.Println(v)
		}
	}
}

func createArt(text string, template []string) {
	var arr []rune
	Newline := false
	for i, r := range text {
		if Newline {
			Newline = false
			printArt(arr, template)
			arr = []rune{}
			continue
		}

		if r == 92 && len(text) != i+1 {
			if text[i+1] == 110 {
				Newline = true
				continue
			}
		}
		arr = append(arr, r)
	}
	printArt(arr, template)
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

func readCmd() (string, string, bool) {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Expected only two input arguments. Try: \"$./ascii-art-fs <text> <option>\"")
		return "", "", true
	}
	if hasError(args[0]) {
		fmt.Println("Input has non readable Ascii characters")
		return "", "", true
	}
	return args[0], args[1], false
}

func hasError(str string) bool {
	for _, r := range str {
		if r < 32 || r > 126 {
			return true
		}
	}
	return false
}

//Printing given rune array, based on lines art
func printArt(arr []rune, lines []string) {
	for line := 1; line <= 8; line++ {
		for _, r := range arr {
			skip := (r - 32) * 9
			fmt.Print(lines[line+int(skip)])
		}
		fmt.Println()
	}
}
