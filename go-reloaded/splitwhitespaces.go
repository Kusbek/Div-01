package piscine

func SplitWhiteSpaces(str string) []string {
	ch := []rune(str)
	j := 0
	l := false

	var res []string
	for i, v := range ch {
		if l && (v == ' ' || v == '\n' || v == '\t') {
			word := string(ch[j:i])
			l = false
			res = MyAppend(res, word)
		}

		if !l && !(v == ' ' || v == '\n' || v == '\t') {
			j = i
			l = true
		}

		if l && i == strlen(str)-1 {
			word := string(ch[j : i+1])
			res = MyAppend(res, word)
		}
	}
	return res

}

func MyAppend(strarray []string, str string) []string {
	len := 0
	for range strarray {
		len++
	}
	newarr := make([]string, len+1)

	for i, v := range strarray {
		newarr[i] = v
	}
	newarr[len] = str

	return newarr
}
