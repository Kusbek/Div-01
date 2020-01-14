package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	//Fetching the argument, and checking for validity
	Arg := os.Args[1:]
	if len(Arg) != 1 {
		return
	}
	for _, r := range Arg[0] {
		if r < 32 || r > 126 {
			return
		}
	}

	//Creating array of strings from the standard
	bytes, err := ioutil.ReadFile("standard.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	lines := strings.Split(string(bytes), "\n")

	//Creating the art itself
	var arr []rune
	Newline := false
	for i, r := range Arg[0] {
		if Newline {
			Newline = false
			printArt(arr, lines)
			arr = []rune{}
			continue
		}

		if r == 92 && len(Arg[0]) != i+1 {
			if Arg[0][i+1] == 110 {
				Newline = true
				continue
			}
		}
		arr = append(arr, r)
	}
	printArt(arr, lines)
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
