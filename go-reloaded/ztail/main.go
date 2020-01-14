package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	argLen := 0
	for range args {
		argLen++
	}
	if argLen == 0 {
		return
	}

	o, n := locate(os.Args)
	option := os.Args[o]
	nbrOfBytes := Atoi(os.Args[n])
	if nbrOfBytes == 0 {
		fmt.Println("invalid number of bytes:", os.Args[n])
		return
	}
	fileNames := os.Args[1:o]
	fileNames = append(fileNames, os.Args[n+1:]...)
	// fmt.Println(option, nbrOfBytes, fileNames)
	len := 0

	if option != "-c" {
		return
	}
	for range fileNames {
		len++
	}
	if len == 0 {
		return
	}
	multiple := false
	if len > 1 {
		multiple = true
	}

	for i, fileName := range fileNames {
		if i != len-1 && i != 0 {
			fmt.Printf("\n")
		}
		read(fileName, nbrOfBytes, multiple)
	}
	os.Exit(0)
}

func locate(args []string) (int, int) {
	for i, v := range args {
		if v == "-c" {
			return i, i + 1
		}
	}
	return 0, 0
}

func read(fileName string, bytes int, multiple bool) {
	file, err := os.Open(fileName)

	// text, _ := ioutil.ReadAll(file)

	if err != nil {
		result := fmt.Sprintf(`ztail: cannot open '%s' for reading: No such file or directory`, fileName)
		fmt.Println(result)
		return
	}

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("%v\n", err.Error())
		return
	}

	fileSize := fileInfo.Size()
	var buffer []byte
	for i := 0; i < int(fileSize); i++ {
		buffer = append(buffer, byte(0))
	}
	bytesread, err := file.Read(buffer)

	if multiple {
		fmt.Printf("==> %v <==", fileName)
		fmt.Printf("\n")
	}
	for _, v := range buffer[(bytesread - bytes):] {
		fmt.Printf("%v", string(v))
	}
}

func mylen(text []byte) int {
	len := 0
	for range text {
		len++
	}
	return len
}
func Atoi(s string) int {
	if s == "" {
		return 0
	}
	var integers []rune = []rune(s)
	var nonzero bool = false
	var nb int = 0
	var length int = StrLen(s)
	var negative bool = false
	var start int = 0
	if integers[0] == '-' {
		negative = true
		start = 1
	} else if integers[0] == '+' {
		start = 1
	}
	for i := start; i < length; i++ {
		if integers[i] < '0' || integers[i] > '9' {
			// fmt.Print("HELLO")
			return 0
		}
		if integers[i] != '0' {
			nonzero = true
		}

		if nonzero {
			nb = nb*10 + int(integers[i]) - 48
			// fmt.Print(int(integer) - 48)
		}
	}

	if negative {
		return -1 * nb
	}
	return nb
}
func StrLen(str string) int {
	var chars []rune = []rune(str)
	var count int = 0
	for i := range chars {
		count = i
	}

	return count + 1
}
