package solver

import (
	"container/list"
)

type Node struct {
	Name      string
	X         int
	Y         int
	Capacity  int
	Used      bool
	Visited   bool
	Neighbors []*Node
}

func Solver(start, end *Node, n map[string]*Node) []string {
	q := list.New()
	q.PushBack(start)
	start.Visited = true
	var mapOfLinks map[string]string = make(map[string]string)
	for q.Len() > 0 {
		e := q.Front()
		node := e.Value.(*Node)
		for _, v := range node.Neighbors {
			if !v.Visited {
				v.Visited = true
				mapOfLinks[v.Name] = node.Name
				q.PushBack(v)
			}
		}
		q.Remove(e)
	}
	path := extractPath(mapOfLinks, n, start, end)
	clearVisited(n)
	return path
}
func clearVisited(n map[string]*Node) {
	for _, v := range n {
		if !v.Used {
			v.Visited = false
		}
	}
}

func extractPath(l map[string]string, n map[string]*Node, start, end *Node) []string {

	name := end.Name
	var result []string
	result = append(result, name)
	ok := false
	for name != start.Name {
		name, ok = l[name]
		if !ok {
			return result
		}
		if name != start.Name && name != end.Name {
			n[name].Used = true
		}
		result = append(result, name)
	}

	return result
}

//         ________________
//        /                \
//  ____[5]----[3]--[1]     |
// /            |    /      |
// [6]---[0]----[4] /       |
// \   ________/|  /        |
//  \ /        [2]/________/
//  [7]_________/
