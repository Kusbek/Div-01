package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	var vowels []rune
	str := Join(args, " ")
	var ch []rune = []rune(str)

	for _, vi := range ch {
		if isVowel(vi) {
			vowels = append(vowels, vi)
		}
	}
	reverse(vowels)

	j := 0
	for i, vi := range ch {
		if isVowel(vi) {
			ch[i] = vowels[j]
			j++
		}
		z01.PrintRune(ch[i])
	}
	z01.PrintRune('\n')
}
func Join(strs []string, sep string) string {
	var result string = ""

	for i, v := range strs {
		if i != 0 {
			result += sep + v
		} else {
			result += v
		}
	}

	return result
}
func reverse(vowels []rune) {
	len := 0
	for range vowels {
		len++
	}

	for i, j := 0, len-1; i < j; i, j = i+1, j-1 {
		vowels[i], vowels[j] = vowels[j], vowels[i]
	}
}
func isVowel(ch rune) bool {
	var vowels []rune = []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	for _, v := range vowels {
		if ch == v {
			return true
		}
	}
	return false
}
