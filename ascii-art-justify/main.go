package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	text, fonttype, alignment, e := readCmd()
	if e {
		return
	}
	template, e := openFileAndReadLineByLine(fonttype)
	if e {
		return
	}
	segmentedTemplate := segment(template)
	createArt(text, segmentedTemplate, alignment)
}
func createArt(text string, template [][]string, alignment string) {
	var arr []string = []string{"", "", "", "", "", "", "", ""}
	Newline := false
	for i, r := range text {
		if Newline {
			Newline = false
			print(arr, alignment)
			arr = []string{"", "", "", "", "", "", "", ""}
			continue
		}
		if r == 92 && len(text) != i+1 {
			if text[i+1] == 110 {
				Newline = true
				continue
			}
		}
		arr = customAppend(arr, template[int(r)-32])
	}
	print(arr, alignment)
}

func print(str []string, alignment string) {
	n := len(str)
	artlen := len(str[0])
	width := getTerminalWidth()
	if alignment == "center" {
		diff := width/2 - artlen/2
		for i := 0; i < n; i++ {
			for j := 0; j < diff; j++ {
				fmt.Print(" ")
			}
			fmt.Println(str[i])
		}
	} else if alignment == "right" {
		diff := width - artlen
		for i := 0; i < n; i++ {
			for j := 0; j < diff; j++ {
				fmt.Print(" ")
			}
			fmt.Println(str[i])
		}
	} else if alignment == "left" {
		for i := 0; i < n; i++ {
			fmt.Println(str[i])
		}
	} else if alignment == "justify" {
		return
	} else {
		fmt.Println("You have entered incorrect alignment option!")
		return
	}

}
func customAppend(str, item []string) []string {
	if str == nil {
		str = make([]string, 8)
	}
	for i := 0; i < 8; i++ {
		str[i] = str[i] + item[i]
	}
	return str
}
func segment(template []string) [][]string {
	var result [][]string
	for i := 0; i < len(template)-1; i = i + 9 {
		temp := template[i+1 : i+9]
		result = append(result, temp)
	}
	return result
}

func readCmd() (string, string, string, bool) {
	args := os.Args[1:]
	if len(args) != 3 {
		fmt.Println("For this exercise you have to enter three input paramters. Try: \"$./ascii-art-fs <text> <font type> --align=<option>\"")
		return "", "", "", true
	}
	if hasError(args[0]) {
		fmt.Println("Input has non readable Ascii characters")
		return "", "", "", true
	}
	if args[2][:8] != "--align=" {
		fmt.Println("You have made a mistake in --align=")
		return "", "", "", true
	}

	option := args[2][8:]
	return args[0], args[1], option, false
}
func getTerminalWidth() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, _ := cmd.Output()
	width, _ := strconv.Atoi(strings.Fields(string(out))[1])
	return width
}
func hasError(str string) bool {
	for _, r := range str {
		if r < 32 || r > 126 {
			return true
		}
	}
	return false
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
