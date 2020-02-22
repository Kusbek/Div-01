package artist

import "strconv"

func parseDate(str string) int64 {
	t := split(str, "-")
	res, _ := strconv.Atoi(t[2])
	return int64(res)
}

func split(str, charset string) []string {
	clen := len(charset)
	var res []string
	l := true
	j := 0
	word := ""
	for i := 0; i < len(str); i++ {
		if i+clen <= len(str) {
			word = str[i : i+clen]
		} else {
			word = str[i:len(str)]
		}
		if !l && word == charset {
			i = i + clen - 1
		} else if !l && word != charset {
			l = true
			j = i
		} else if l && word == charset {
			// fmt.Println(str[j:i])
			res = append(res, str[j:i])
			i = i + clen - 1
			l = false
		} else if l && i == len(str)-1 {
			word := string(str[j : i+1])
			res = append(res, word)
		}
	}

	return res
}

// 		/////////////////////////////////////////////////////////////////
// 		locations := getLocations(strconv.Itoa(int(artist.ID))).Locations
// 		var countries []string
// 		for _, location := range locations {
// 			countries = append(countries, split(location, "-")[1])
// 		}
// 		for _, country := range countries {
// 			if !set[country] {
// 				set[country] = true
// 			}
// 		}
// 		fmt.Println(countries)

// 		/////////////////////////////////////////////////////////////////

// for k := range set { // Loop
// 	fmt.Println(k)
// }
