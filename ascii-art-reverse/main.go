package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	arg := os.Args[1:]
	if len(arg) < 1 {
		return
	}
	if !strings.Contains(arg[0], "--reverse=") {
		return
	}
	template, e := openFileAndReadLineByLine("standard.txt")
	if e {
		return
	}
	segmentedTemplate := segment(template)
	targetTemplate, e := openFileAndReadLineByLine(arg[0][10:])
	if e {
		return
	}

	str := ""
	for len(targetTemplate[0]) > 0 {
		for i, v := range segmentedTemplate {
			if checkIfLetterIsPresent(v, targetTemplate) {
				str = str + string(rune(i+32))
				targetTemplate = removeLetter(len(v[0]), targetTemplate)
			}
		}
	}
	fmt.Println(str)
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

func openFileAndReadLineByLine(o string) ([]string, bool) {
	file := fmt.Sprintf(o)
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return nil, true
	}
	lines := strings.Split(string(bytes), "\n")

	return lines, false
}
