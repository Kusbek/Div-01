package piscine

import (
	"github.com/01-edu/z01"
)

func PrintCombN(n int) {
	var i1 rune
	var first bool = true
	for i1 = 48; i1 < 58; i1++ {
		if n == 1 {
			if first == true {
				first = false
			} else {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			z01.PrintRune(i1)
			continue
		}
		for i2 := i1 + 1; i2 < 58; i2++ {
			if n == 2 {
				if first == true {
					first = false
				} else {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
				z01.PrintRune(i1)
				z01.PrintRune(i2)
				continue
			}
			for i3 := i2 + 1; i3 < 58; i3++ {
				if n == 3 {
					if first == true {
						first = false
					} else {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
					z01.PrintRune(i1)
					z01.PrintRune(i2)
					z01.PrintRune(i3)
					continue
				}
				for i4 := i3 + 1; i4 < 58; i4++ {
					if n == 4 {
						if first == true {
							first = false
						} else {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
						z01.PrintRune(i1)
						z01.PrintRune(i2)
						z01.PrintRune(i3)
						z01.PrintRune(i4)
						continue
					}
					for i5 := i4 + 1; i5 < 58; i5++ {
						if n == 5 {
							if first == true {
								first = false
							} else {
								z01.PrintRune(',')
								z01.PrintRune(' ')
							}
							z01.PrintRune(i1)
							z01.PrintRune(i2)
							z01.PrintRune(i3)
							z01.PrintRune(i4)
							z01.PrintRune(i5)
							continue
						}
						for i6 := i5 + 1; i6 < 58; i6++ {
							if n == 6 {
								if first == true {
									first = false
								} else {
									z01.PrintRune(',')
									z01.PrintRune(' ')
								}
								z01.PrintRune(i1)
								z01.PrintRune(i2)
								z01.PrintRune(i3)
								z01.PrintRune(i4)
								z01.PrintRune(i5)
								z01.PrintRune(i6)
								continue
							}
							for i7 := i6 + 1; i7 < 58; i7++ {
								if n == 7 {
									if first == true {
										first = false
									} else {
										z01.PrintRune(',')
										z01.PrintRune(' ')
									}
									z01.PrintRune(i1)
									z01.PrintRune(i2)
									z01.PrintRune(i3)
									z01.PrintRune(i4)
									z01.PrintRune(i5)
									z01.PrintRune(i6)
									z01.PrintRune(i7)
									continue
								}
								for i8 := i7 + 1; i8 < 58; i8++ {
									if n == 8 {
										if first == true {
											first = false
										} else {
											z01.PrintRune(',')
											z01.PrintRune(' ')
										}
										z01.PrintRune(i1)
										z01.PrintRune(i2)
										z01.PrintRune(i3)
										z01.PrintRune(i4)
										z01.PrintRune(i5)
										z01.PrintRune(i6)
										z01.PrintRune(i7)
										z01.PrintRune(i8)
										continue
									}
									for i9 := i8 + 1; i9 < 58; i9++ {
										if n == 9 {
											if first == true {
												first = false
											} else {
												z01.PrintRune(',')
												z01.PrintRune(' ')
											}
											z01.PrintRune(i1)
											z01.PrintRune(i2)
											z01.PrintRune(i3)
											z01.PrintRune(i4)
											z01.PrintRune(i5)
											z01.PrintRune(i6)
											z01.PrintRune(i7)
											z01.PrintRune(i8)
											z01.PrintRune(i9)
											continue
										}
									}
								}

							}
						}

					}

				}

			}

		}

	}
	z01.PrintRune('\n')
}
