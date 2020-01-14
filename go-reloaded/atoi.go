package piscine

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
func mylen(s []rune) int {
	len := 0
	for range s {
		len++
	}
	return len
}
