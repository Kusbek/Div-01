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
	segmentedTemplate := segment(template)
	targetTemplate, e := openFileAndReadLineByLine("file")
	if e {
		return
	}

	str := ""
	for len(targetTemplate[0]) > 0 {
		for i, v := range segmentedTemplate {
			if checkIfLetterIsPresent(v, targetTemplate) {
				str = str + string(rune(i+32))
				targetTemplate = removeLetter(len(v[0]), targetTemplate)
				fmt.Println(str, len(v[0]))
				for _, v := range targetTemplate {
					fmt.Println(v)
				}
			}
		}
	}
}

func removeLetter(length int, word []string) []string {
	for i, v := range word[0 : len(word)-1] {
		// fmt.Println(v[length:])
		word[i] = v[length:]
	}
	return word
}
func checkIfLetterIsPresent(letter, word []string) bool {
	found := true
	if len(letter[0]) > len(word[0]) {
		return false
	}
	for i, v := range word[0 : len(word)-1] {
		if letter[i] != v[:len(letter[i])] {
			found = false
		}
	}
	return found
}
func segment(template []string) [][]string {
	var result [][]string
	for i := 0; i < len(template)-1; i = i + 9 {
		temp := template[i+1 : i+9]
		result = append(result, temp)
	}
	return result
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
