package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	if hasError(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	str := ""
	b := []rune(base)
	bl := strlen(base)
	if nbr < 0 {
		if nbr == -9223372036854775808 {
			nbr = (nbr + 1) * -1
			rem := nbr % bl
			nbr = nbr / bl
			z01.PrintRune('-')
			str = string(b[rem+1]) + str
		} else {
			z01.PrintRune('-')
			nbr = -nbr
		}
	}

	for nbr > 0 {
		rem := nbr % bl
		nbr = nbr / bl
		str = string(b[rem]) + str
	}

	for _, v := range str {
		z01.PrintRune(v)
	}
}
func strlen(base string) int {
	len := 0
	for range base {
		len++
	}
	return len
}
func hasError(base string) bool {
	b := []rune(base)
	length := 0
	for i, v := range b {
		length = i + 1
		if v == '+' || v == '-' {
			return true
		}

		for j := 0; j < i; j++ {
			if b[i] == b[j] {
				return true
			}
		}
	}
	if length < 2 {
		return true
	}
	return false
}
