package piscine

func AtoiBase(s string, base string) int {
	bsl := strlen(base)
	nlen := strlen(s)
	vals := make([]int, nlen)

	for i, v := range s {
		for ii, vv := range base {
			if v == vv {
				vals[i] = ii
			}
		}
	}
	if hasError(base) {
		return 0
	}
	n := nlen - 1
	res := 0
	for _, v := range vals {
		res += v * IterativePower(bsl, n)
		n--
	}

	return res

}

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	var res int = 1
	for i := 0; i < power; i++ {
		res *= nb
	}
	return res
}
