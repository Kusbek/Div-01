package piscine

func AdvancedSortWordArr(array []string, f func(a, b string) int) {
	for i := 0; i < len(array); i++ {
		for ii := i + 1; ii < len(array); ii++ {
			out := f(array[i], array[ii])
			if out == 1 {
				array[i], array[ii] = array[ii], array[i]
			}
		}
	}
}

func Compare(a, b string) int {
	if a == b {
		return 0
	} else if a < b {
		return -1
	}

	return 1
}
