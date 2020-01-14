package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	l := 0
	for range args {
		l++
	}
	if l != 3 {
		fmt.Println(0)
		return
	}
	n1 := args[0]
	n2 := args[2]
	o := args[1]

	if !IsNumeric(n1) || !IsNumeric(n2) {
		fmt.Println(0)
		return
	}

	num1, e1 := myAtoi(n1)
	num2, e2 := myAtoi(n2)

	if e1 || e2 {
		fmt.Println(0)
		return
	}

	result, e := doMath(num1, num2, o)

	if e {
		fmt.Println(0)
		return
	}

	fmt.Println(result)

}
func doMath(n1, n2 int, op string) (int, bool) {
	switch op {
	case "+":
		if (n1 > 0 && n2 > 0 && n1+n2 < 0) || (n1 < 0 && n2 < 0 && n1+n2 > 0) {
			fmt.Println("Overflow")
			return 0, true
		}
		return n1 + n2, false
	case "-":
		if (n1 > 0 && n2 < 0 && n1+n2 < 0) || (n1 < 0 && n2 > 0 && n1+n2 > 0) {
			fmt.Println("Overflow")
			return 0, true
		}
		return n1 - n2, false
	case "*":
		result := n1 * n2
		if n2 != result/n1 {
			fmt.Println("Overflow")
			return 0, true
		}
		return result, false
	case "/":
		if n2 == 0 {
			fmt.Print("No division by ")
			return 0, false
		}
		return n1 / n2, false
	case "%":
		if n2 == 0 {
			fmt.Print("No Modulo by ")
			return 0, false
		}
		return n1 % n2, false
	}
	return 0, false
}
func myAtoi(s string) (int, bool) {
	ch := []rune(s)
	start := 0
	neg := false
	if ch[0] == '+' || ch[0] == '-' {
		start = 1
	}
	if ch[0] == '-' {
		neg = true
	}
	result := 0
	for i := start; i < mylen(ch); i++ {
		if !(ch[i] >= '0' && ch[i] < '9') {
			return 0, true
		}
		if neg {
			result = result*10 - (int(ch[i]) - 48)
			if result > 0 {
				return 0, true
			}
		} else {
			result = result*10 + (int(ch[i]) - 48)
			if result < 0 {
				return 0, true
			}
		}
	}
	return result, false
}

func mylen(s []rune) int {
	len := 0
	for range s {
		len++
	}
	return len
}

func IsNumeric(str string) bool {
	var ch []rune = []rune(str)
	for i := range ch {
		if !(ch[i] >= '0' && ch[i] <= '9') && ch[i] != '-' {
			return false
		}
	}
	return true
}
