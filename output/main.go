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
	Num := len(Arg)
	if Num < 1 {
		return
	}
	output := false
	outputTtxt := ""
	style := "standard"
	//Creating array of strings from the standard

	if strings.Contains(Arg[0], "--reverse=") {
		return
	}

	if Num > 1 {
		style = Arg[1]
	}

	if Num > 2 {
		for _, arg := range Arg[2:] {
			for _, r := range arg {
				if r < 32 || r > 126 {
					return
				}
			}
			if strings.Contains(arg, "--output=") {
				output = true
				for i := 9; i != len(arg); i++ {
					outputTtxt += string(arg[i])
				}
			}
		}
	}
	art := ""
	var lines []string

	bytes, err := ioutil.ReadFile(style)
	if err != nil {
		fmt.Println(err)
		return
	}
	lines = strings.Split(string(bytes), "\n")

	art = asciiart(Arg[0], lines)
	if output {
		err2 := ioutil.WriteFile(outputTtxt, []byte(art), 0644)
		check(err2)
	} else {
		fmt.Println(art)
	}
}
func check(e error) {
	if e != nil {
		panic(e)
	}
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

//Printing given rune array, based on lines art
func printArt(arr []rune, lines []string) {
	for line := 0; line < 8; line++ {
		for _, r := range arr {
			skip := (r-32)*9 + 1
			fmt.Print(lines[line+int(skip)])
		}
		fmt.Println()
	}
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
