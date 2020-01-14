package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	var len int = 0

	for range args {
		len++
	}

	if len == 0 {
		io.Copy(os.Stdout, os.Stdin)
		return
	}

	for _, fileName := range args {
		file, err := os.Open(fileName)
		text, _ := ioutil.ReadAll(file)
		if err != nil {
			PrintStr(err.Error())
		} else {
			PrintStr(string(text))
		}
		// z01.PrintRune('\n')

		// if i != len-1 && err == nil {
		// 	z01.PrintRune('\n')
		// }
	}
}

func PrintStr(str string) {
	var s []rune = []rune(str)
	for _, v := range s {
		z01.PrintRune(v)
	}
}
