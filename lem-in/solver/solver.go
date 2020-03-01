package solver

import "container/list"

type Node struct {
	Name      string
	X         int
	Y         int
	Capacity  int
	Visited   bool
	Neighbors []*Node
}

func Solver(numberOfAnts int, start, end *Node) {
	q := list.New()
	q.PushBack(start)

	for q.Len() > 0 {

	}
}

//         ________________
//        /                \
//  ____[5]----[3]--[1]     |
// /            |    /      |
// [6]---[0]----[4] /       |
// \   ________/|  /        |
//  \ /        [2]/________/
//  [7]_________/
