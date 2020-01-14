package piscine

func Split(str, charset string) []string {
	clen := strlen(charset)
	var res []string
	l := true
	j := 0
	word := ""
	for i := 0; i < strlen(str); i++ {
		if i+clen <= strlen(str) {
			word = str[i : i+clen]
		} else {
			word = str[i:strlen(str)]
		}
		if !l && word == charset {
			i = i + clen - 1
		} else if !l && word != charset {
			l = true
			j = i
		} else if l && word == charset {
			// fmt.Println(str[j:i])
			res = MyAppend(res, str[j:i])
			i = i + clen - 1
			l = false
		} else if l && i == strlen(str)-1 {
			word := string(str[j : i+1])
			res = MyAppend(res, word)
		}
	}

	return res
}
