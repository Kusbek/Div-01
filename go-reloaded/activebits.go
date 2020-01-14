package piscine

func ActiveBits(n int) uint {
	var count uint = 0
	for n > 0 {
		rem := n % 2
		n = n / 2
		if rem == 1 {
			count++
		}
	}
	return count
}
