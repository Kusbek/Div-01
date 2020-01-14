package main

import (
	"fmt"

	student ".."
	// "github.com/01-edu/z01"
)

func main() {

	// //Atoi
	// fmt.Println("Atoi")
	// s := ""
	// s2 := "-"
	// s3 := "--123"
	// s4 := "1"
	// s5 := "-3"
	// s6 := "8292"
	// s7 := "9223372036854775807"
	// s8 := "-9223372036854775808"
	// n := student.Atoi(s)
	// n2 := student.Atoi(s2)
	// n3 := student.Atoi(s3)
	// n4 := student.Atoi(s4)
	// n5 := student.Atoi(s5)
	// n6 := student.Atoi(s6)
	// n7 := student.Atoi(s7)
	// n8 := student.Atoi(s8)
	// fmt.Println(n)
	// fmt.Println(n2)
	// fmt.Println(n3)
	// fmt.Println(n4)
	// fmt.Println(n5)
	// fmt.Println(n6)
	// fmt.Println(n7)
	// fmt.Println(n8)

	// // //Recursive Power
	// fmt.Println("Recursive Power")
	// fmt.Println(student.RecursivePower(-7, -2))
	// fmt.Println(student.RecursivePower(-8, -7))
	// fmt.Println(student.RecursivePower(4, 8))
	// fmt.Println(student.RecursivePower(1, 3))
	// fmt.Println(student.RecursivePower(-1, 1))
	// fmt.Println(student.RecursivePower(-6, 5))

	// // //PrintCombN
	// fmt.Println("PrintCombN")
	// for i := 1; i < 10; i++ {
	// 	student.PrintCombN(i)
	// 	fmt.Println("End")
	// }

	// //PrintNbrBase
	// fmt.Println("PrintNbrBase")
	// student.PrintNbrBase(919617, "01")
	// z01.PrintRune('\n')
	// student.PrintNbrBase(753639, "CHOUMIisDAcat!")
	// z01.PrintRune('\n')
	// student.PrintNbrBase(-174336, "CHOUMIisDAcat!")
	// z01.PrintRune('\n')
	// student.PrintNbrBase(-861737, "Zone01Zone01!")
	// z01.PrintRune('\n')
	// student.PrintNbrBase(-661165, "1")
	// z01.PrintRune('\n')
	// student.PrintNbrBase(125, "0123456789ABCDEF")
	// z01.PrintRune('\n')
	// student.PrintNbrBase(-125, "choumi")
	// z01.PrintRune('\n')
	// student.PrintNbrBase(125, "-ab")
	// z01.PrintRune('\n')
	// student.PrintNbrBase(-9223372036854775808, "0123456789")
	// z01.PrintRune('\n')

	// //Ato"fmt"iBase
	// fmt.Println("AtoiBase")
	// fmt.Println(student.AtoiBase("bcbbbbaab", "abc"))
	// fmt.Println(student.AtoiBase("0001", "01"))
	// fmt.Println(student.AtoiBase("00", "01"))
	// fmt.Println(student.AtoiBase("saDt!I!sI", "CHOUMIisDAcat!"))
	// fmt.Println(student.AtoiBase("AAho?Ao", "WhoAmI?"))
	// fmt.Println(student.AtoiBase("thisinputshouldnotmatter", "abca"))
	// fmt.Println(student.AtoiBase("125", "0123456789"))
	// fmt.Println(student.AtoiBase("uoi", "choumi"))
	// fmt.Println(student.AtoiBase("bbbbbab", "-ab"))

	// //splitwhitespaces
	// fmt.Println(student.SplitWhiteSpaces("The earliest foundations of what would become computer science predate the invention of the modern digital computer"))
	// fmt.Println(student.SplitWhiteSpaces("Machines for calculating fixed numerical tasks such as the abacus have existed since antiquity,"))
	// fmt.Println(student.SplitWhiteSpaces("aiding in computations such as multiplication and division ."))
	// fmt.Println(student.SplitWhiteSpaces("Algorithms for performing computations have existed since antiquity, even before the development of sophisticated computing equipment."))
	// fmt.Println(student.SplitWhiteSpaces("Wilhelm Schickard designed and constructed the first working mechanical calculator in 1623.[4]"))
	// fmt.Println(student.SplitWhiteSpaces(" In 1673, Gottfried Leibniz demonstrated a digital mechanical calculator,"))

	// //split
	// str := []string{
	// 	"|=choumi=|which|=choumi=|itself|=choumi=|used|=choumi=|cards|=choumi=|and|=choumi=|a|=choumi=|central|=choumi=|computing|=choumi=|unit.|=choumi=|When|=choumi=|the|=choumi=|machine|=choumi=|was|=choumi=|finished,",
	// 	"!==!which!==!was!==!making!==!all!==!kinds!==!of!==!punched!==!card!==!equipment!==!and!==!was!==!also!==!in!==!the!==!calculator!==!business[10]!==!to!==!develop!==!his!==!giant!==!programmable!==!calculator,",
	// 	"AFJCharlesAFJBabbageAFJstartedAFJtheAFJdesignAFJofAFJtheAFJfirstAFJautomaticAFJmechanicalAFJcalculator,",
	// 	"<<==123==>>In<<==123==>>1820,<<==123==>>Thomas<<==123==>>de<<==123==>>Colmar<<==123==>>launched<<==123==>>the<<==123==>>mechanical<<==123==>>calculator<<==123==>>industry[note<<==123==>>1]<<==123==>>when<<==123==>>he<<==123==>>released<<==123==>>his<<==123==>>simplified<<==123==>>arithmometer,"}
	// charset := []string{
	// 	"|=choumi=|",
	// 	"!==!",
	// 	"AFJ",
	// 	"<<==123==>>"}
	// fmt.Println("split")
	// for i := range str {
	// 	fmt.Println(student.Split(str[i], charset[i]))
	// }

	// // //Convertbase
	// fmt.Println("Convert base")
	// fmt.Println(student.ConvertBase("4506C", "0123456789ABCDEF", "choumi"))
	// fmt.Println(student.ConvertBase("babcbaaaaabcb", "abc", "0123456789ABCDEF"))
	// fmt.Println(student.ConvertBase("256850", "0123456789", "01"))
	// fmt.Println(student.ConvertBase("uuhoumo", "choumi", "Zone01"))
	// fmt.Println(student.ConvertBase("683241", "0123456789", "0123456789"))

	// // //AdvancedSortWordArr
	// a := []string{"The", "earliest", "computing", "device", "undoubtedly", "consisted", "of", "the", "five", "fingers", "of", "each", "hand"}
	// student.AdvancedSortWordArr(a, student.Compare)
	// fmt.Println(a)
	// d := []string{"The", "word", "digital", "comesfrom", "\"digits\"", "or", "fingers"}
	// student.AdvancedSortWordArr(d, student.Compare)
	// fmt.Println(d)
	// b := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	// student.AdvancedSortWordArr(b, student.Compare)
	// fmt.Println(b)
	// e := []string{"The", "computing", "consisted", "device", "each", "earliest", "fingers", "five", "hand", "of", "of", "the", "undoubtedly"}
	// student.AdvancedSortWordArr(e, func(a, b string) int { return student.Compare(b, a) })
	// fmt.Println(e)
	// c := []string{"\"digits\"", "The", "comesfrom", "digital", "fingers", "or", "word"}
	// student.AdvancedSortWordArr(c, func(a, b string) int { return student.Compare(b, a) })
	// fmt.Println(c)
	// f := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	// student.AdvancedSortWordArr(f, func(a, b string) int { return student.Compare(b, a) })
	// fmt.Println(f)

	// //ActiveBits
	// fmt.Println("ActiveBits")
	// fmt.Println(student.ActiveBits(15))
	// fmt.Println(student.ActiveBits(17))
	// fmt.Println(student.ActiveBits(4))
	// fmt.Println(student.ActiveBits(11))
	// fmt.Println(student.ActiveBits(9))
	// fmt.Println(student.ActiveBits(12))
	// fmt.Println(student.ActiveBits(2))

	// // //SortPrintList
	// s0 := []int{0}
	// checkSortedListInsert(s0, 39)
	// s1 := []int{0, 1, 2, 3, 4, 5, 24, 25, 54}
	// checkSortedListInsert(s1, 33)
	// s2 := []int{0, 2, 18, 33, 37, 37, 39, 52, 53, 57}
	// checkSortedListInsert(s2, 53)
	// s3 := []int{0, 5, 18, 24, 28, 35, 42, 45, 52}
	// checkSortedListInsert(s3, 52)
	// s4 := []int{0, 12, 20, 23, 23, 24, 30, 41, 53, 57, 59}
	// checkSortedListInsert(s4, 38)

	// //cSortedListMerge
	// s1 := []int{}
	// s2 := []int{}
	// checkSortedListMerge(s1, s2)
	// s3 := []int{}
	// s4 := []int{2, 2, 4, 9, 12, 12, 19, 20}
	// checkSortedListMerge(s3, s4)
	// s5 := []int{}
	// s6 := []int{4, 4, 6, 9, 13, 18, 20, 20}
	// checkSortedListMerge(s5, s6)
	// s7 := []int{0, 7, 39, 92, 97, 93, 91, 28, 64}
	// s8 := []int{80, 23, 27, 30, 85, 81, 75, 70}
	// checkSortedListMerge(s7, s8)
	s11 := []int{1, 2, 11, 30, 54, 56, 70, 79, 99}
	s12 := []int{1, 28, 38, 67, 67, 79, 95, 97}
	checkSortedListMerge(s11, s12)
	// s9 := []int{0, 3, 8, 8, 13, 19, 34, 38, 46}
	// s10 := []int{7, 39, 45, 53, 59, 70, 76, 79}
	// checkSortedListMerge(s9, s10)

	// //ListRemoveIf\
	// strings1 := []string{}
	// ints1 := []int{}
	// dataref1 := 1
	// checkListRemoveIf(ints1, strings1, dataref1)
	// string2 := []string{}
	// ints2 := []int{}
	// dataref2 := 96
	// checkListRemoveIf(ints2, string2, dataref2)

	// str3 := []string{}
	// i3 := []int{98, 98, 33, 34, 33, 34, 33, 89, 33}
	// d3 := 34
	// checkListRemoveIf(i3, str3, d3)

	// s4 := []string{}
	// i4 := []int{79, 74, 99, 79, 7}
	// d4 := 99
	// checkListRemoveIf(i4, s4, d4)

	// s5 := []string{}
	// i5 := []int{56, 93, 68, 56, 87, 68, 56, 68}
	// d5 := 68
	// checkListRemoveIf(i5, s5, d5)

	// strings := []string{"mvkUxbqhQve4l", "4Zc4t hnf SQ", "q2If E8BPuX"}
	// ints := []int{}
	// dataref := "4Zc4t hnf SQ"
	// checkListRemoveIf(ints, strings, dataref)

	// //BTreeTransplant
	// root := "01"
	// strings := []string{"07", "12", "05", "10", "02", "03"}
	// node := "12"
	// root1 := "55"
	// strings1 := []string{"60", "33", "12", "15"}
	// checkBTreeTransplant(root, strings, node, root1, strings1)

	// r2 := "03"
	// s2 := []string{"39", "99", "11", "44", "14", "11"}
	// n2 := "11"
	// r12 := "55"
	// s12 := []string{"60", "33", "12", "15"}
	// checkBTreeTransplant(r2, s2, n2, r12, s12)

	// r := "33"
	// s := []string{"05", "20", "52", "31", "13", "11"}
	// n := "20"
	// r1 := "55"
	// s1 := []string{"60", "33", "12", "15"}
	// checkBTreeTransplant(r, s, n, r1, s1)

	// //BTreeApplyByLevel
	// root := "01"
	// strings := []string{"07", "12", "05", "10", "02", "03"}
	// checkBTreeApplyByLevel(root, strings)
	// root := "01"
	// strings := []string{"07", "12", "05", "10", "02", "03"}
	// checkBTreeApplyByLevel(root, strings)

	// BTreeDeleteNode
	// r := "01"
	// s := []string{"07", "12", "05", "10", "02", "03"}
	// d := "02"
	// checkBTreeDeleteNode(r, s, d)
	// root := "33"
	// strings := []string{"5", "20", "31", "13", "52", "11"}
	// delete := "20"
	// checkBTreeDeleteNode(root, strings, delete)

	// r1 := "03"
	// s1 := []string{"39", "99", "11", "44", "14", "11"}
	// d1 := "03"
	// checkBTreeDeleteNode(r1, s1, d1)

	// r2 := "03"
	// s2 := []string{"03", "01", "94", "19", "24", "111"}
	// d2 := "03"
	// checkBTreeDeleteNode(r2, s2, d2)

}

func checkBTreeDeleteNode(myroot string, strings []string, delete string) {
	root := &student.TreeNode{Data: myroot}
	for _, v := range strings {
		student.BTreeInsertData(root, v)
	}
	node := student.BTreeSearchItem(root, delete)
	root = student.BTreeDeleteNode(root, node)
	student.BTreeApplyInorder(root, fmt.Println)
}
func checkBTreeApplyByLevel(myroot string, strings []string) {
	root := &student.TreeNode{Data: myroot}
	for _, v := range strings {
		student.BTreeInsertData(root, v)
	}
	student.BTreeApplyByLevel(root, fmt.Print)
}

func checkBTreeTransplant(rootValue string, branches []string, rplc string, rplcroot string, rplcbranches []string) {
	root := &student.TreeNode{Data: rootValue}
	for _, v := range branches {
		student.BTreeInsertData(root, v)
	}

	replacement := &student.TreeNode{Data: rplcroot}

	for _, v := range rplcbranches {
		student.BTreeInsertData(replacement, v)
	}

	node := student.BTreeSearchItem(root, rplc)
	root = student.BTreeTransplant(root, node, replacement)
	student.BTreeApplyInorder(root, fmt.Println)
}
func checkListRemoveIf(ints []int, start []string, dataref interface{}) {
	var link *student.List = &student.List{}
	for _, v := range start {
		student.ListPushBack(link, v)
	}

	for _, v := range ints {
		student.ListPushBack(link, v)
	}
	student.ListRemoveIf(link, dataref)
	PrintList1(link)
}
func checkSortedListMerge(n1, n2 []int) {
	var link1 *student.NodeI
	var link2 *student.NodeI

	for _, v := range n1 {
		link1 = listPushBack(link1, v)
	}

	for _, v := range n2 {
		link2 = listPushBack(link2, v)
	}

	link := student.SortedListMerge(link1, link2)
	PrintList(link)
}

func checkSortedListInsert(start []int, data_ref int) {
	var link *student.NodeI
	for _, v := range start {
		link = listPushBack(link, v)
	}

	link = student.SortListInsert(link, data_ref)
	PrintList(link)
}

func PrintList(l *student.NodeI) {
	it := l
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}

func PrintList1(l *student.List) {
	it := l.Head
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}

	fmt.Print(nil, "\n")
}

func listPushBack(l *student.NodeI, data int) *student.NodeI {
	n := &student.NodeI{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}
