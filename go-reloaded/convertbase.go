package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	if nbr == "0" {
		return "0"
	}

	basefrommap := []rune(baseFrom)   // 0123456789ABCDEF
	basefromint := mylen(basefrommap) //16
	number := []rune(nbr)             //FFFF
	nbrlen := mylen(number)           // 4
	tenbase := 0

	for _, v := range number {
		n := 0
		for ii, vv := range basefrommap {
			if v == vv {
				n = ii
			}
		}

		nbrlen--
		tenbase += n * IterativePower(basefromint, nbrlen)
	}
	result := mufunc(tenbase, baseTo)

	return result
}

func mufunc(nbr int, base string) string {
	runebase := []rune(base)
	intbase := mylen(runebase)
	return convertToBase(nbr, intbase, runebase)
}

func convertToBase(nbr, base int, runebase []rune) string {
	if nbr == 0 {
		return ""
	}
	rem := nbr % base
	nbr = nbr / base
	result := convertToBase(nbr, base, runebase)
	return result + string(runebase[rem])
}
