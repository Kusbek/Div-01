package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	s "DIV-01/lem-in/solver"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}
	filename := args[0]

	numberOfAnts, lines := readLines(filename)
	if numberOfAnts < 1 {
		var err error = fmt.Errorf("number of ants is invalid")
		abortOnError(err)
	}
	start := time.Now()
	mapOfNodes, startNode, endNode := parseLines(lines)

	for _, v := range mapOfNodes[startNode.Name].Neighbors {
		if endNode == v {
			for i := 1; i <= numberOfAnts; i++ {
				fmt.Printf("L%d-%s ", i, endNode.Name)
			}
			fmt.Println()
			return
		}
	}
	var takenPath []string
	allPaths := Solver(takenPath, mapOfNodes, startNode, endNode)
	paths := allocatePaths(allPaths, startNode, mapOfNodes)
	for i, path := range paths {
		paths[i] = path[1:]
	}
	lemin(numberOfAnts, paths, mapOfNodes)
	elapsed := time.Since(start)
	log.Printf("took %s", elapsed)
}

func allocatePaths(paths [][]string, startNode *s.Node, mapOfNodes map[string]*s.Node) [][]string {
	var maxNPaths [][][]string = make([][][]string, len(paths))
	for i := 0; i < len(paths); i++ {
		maxNPaths[i] = append(maxNPaths[i], paths[i])
		for j := i + 1; j < len(paths); j++ {
			if norm(maxNPaths[i], paths[j]) {
				maxNPaths[i] = append(maxNPaths[i], paths[j])
			}
		}
	}
	var result [][][]string
	max := 0
	for _, v := range maxNPaths {
		if max < len(v) {
			max = len(v)
		}
	}

	for _, v := range maxNPaths {
		if max == len(v) {
			result = append(result, v)
		}
	}
	var res [][]string
	min := int(^uint(0) >> 1)
	for _, v := range result {
		tempmin := 0
		for _, vv := range v {
			tempmin += len(vv)
		}

		if tempmin < min {
			min = tempmin
			res = v
		}
	}

	return res
}

func norm(n1 [][]string, n2 []string) bool {
	for _, v := range n1 {
		for _, k := range v[1 : len(v)-1] {
			for _, kk := range n2[1 : len(n2)-1] {
				if k == kk {
					return false
				}
			}
		}
	}
	return true
}

func Solver(takenPath []string, mapOfNodes map[string]*s.Node, startNode, endNode *s.Node) [][]string {
	if startNode == endNode {
		return nil
	}
	extractedPath := reverse(s.BFS(startNode, endNode, mapOfNodes))
	var fullpath []string
	if extractedPath != nil {
		fullpath = append(fullpath, takenPath...)
		fullpath = append(fullpath, extractedPath...)
	}
	takenPath = append(takenPath, startNode.Name)

	startNode.Used = true
	s.ClearVisited(mapOfNodes)
	var result [][]string
	if fullpath != nil {
		result = append(result, fullpath)
	}
	for _, v := range startNode.Neighbors {
		if !v.Used {
			paths := Solver(takenPath, mapOfNodes, v, endNode)
			for _, path := range paths {
				if !isEqual(fullpath, path) {
					result = append(result, path)
				}
			}
		}
	}

	startNode.Used = false
	startNode.Visited = false
	return result
}

func isEqual(fullpath, path []string) bool {
	if len(fullpath) != len(path) {
		return false
	}
	for i := range fullpath {
		if fullpath[i] != path[i] {
			return false
		}
	}

	return true
}
func minLen(p [][]string, ants [][]int) int {
	min := int(^uint(0) >> 1)
	for i := range p {
		if min > len(p[i])+len(ants[i]) {
			min = len(p[i]) + len(ants[i])
		}
	}
	return min
}

func MaxLen(p [][]string, ants [][]int) int {
	min := 0
	for i := range p {
		if min < len(p[i])+len(ants[i]) {
			min = len(p[i]) + len(ants[i])
		}
	}
	return min
}

func lemin(n int, p [][]string, m map[string]*s.Node) {
	var antQueues [][]int = make([][]int, len(p)) //list.List = make([]list.List, len(p))
	i := 1
	min := minLen(p, antQueues)

	for i <= n {
		for k := 0; k < len(p); k++ {
			if len(p[k])+len(antQueues[k]) <= min {
				antQueues[k] = append(antQueues[k], i)
				min = minLen(p, antQueues)
				break
			}
		}

		i++
	}

	max := MaxLen(p, antQueues)
	var solution [][]string = make([][]string, max-1)
	for i := 0; i < len(p); i++ {

		for j, v := range antQueues[i] {
			for k, w := range p[i] {
				str := fmt.Sprintf("L%d-%s", v, w)
				solution[k+j] = append(solution[k+j], str)
			}
		}
	}

	for _, v := range solution {
		for _, w := range v {
			fmt.Printf("%s ", w)
		}
		fmt.Println()
	}
}
func reverse(nodes []string) []string {
	for i, j := 0, len(nodes)-1; i < j; i, j = i+1, j-1 {
		nodes[i], nodes[j] = nodes[j], nodes[i]
	}
	return nodes
}
func abortOnError(err error) {
	if err != nil {
		fmt.Println(err)
		fmt.Println("Normalno dannye vvodi, sumelek")
		os.Exit(1)
	}
}

func buildNodes(nodes [][]string) map[string]*s.Node {
	var n map[string]*s.Node = make(map[string]*s.Node)
	for _, v := range nodes {
		x, err := strconv.Atoi(v[1])
		abortOnError(err)
		y, err := strconv.Atoi(v[2])
		abortOnError(err)
		if _, ok := n[v[0]]; ok {
			var err error = fmt.Errorf("duplicate rooms")
			abortOnError(err)
		}
		n[v[0]] = &s.Node{Name: v[0], X: x, Y: y}
	}
	return n
}

func parseLines(lines []string) (map[string]*s.Node, *s.Node, *s.Node) {
	var start bool
	var end bool
	var startNode string
	var endNode string
	var nodes [][]string
	var links [][]string
	for _, v := range lines {
		if !strings.HasPrefix(v, "#") {
			splitted := strings.Split(v, " ")
			if len(splitted) == 3 {
				nodes = append(nodes, splitted)
				if start {
					startNode = splitted[0]
					start = false
				}
				if end {
					endNode = splitted[0]
					end = false
				}
			} else {
				splitted = strings.Split(v, "-")
				links = append(links, splitted)
			}
		}

		if v == "##start" {
			start = true
		}

		if v == "##end" {
			end = true
		}
	}
	if startNode == "" || endNode == "" {
		var err error = fmt.Errorf("no start or end")
		abortOnError(err)
	}
	n := buildNodes(nodes)
	createLinks(n, links)

	return n, n[startNode], n[endNode]

}

func createLinks(n map[string]*s.Node, links [][]string) {
	for _, link := range links {
		node1, ok := n[link[0]]
		if !ok {
			var err error = fmt.Errorf("unknown room")
			abortOnError(err)
		}
		node2, ok := n[link[1]]
		if !ok {
			var err error = fmt.Errorf("unknown room")
			abortOnError(err)
		}
		if node1 == node2 {
			var err error = fmt.Errorf("self linkage")
			abortOnError(err)
		}
		node1.Neighbors = append(node1.Neighbors, node2)
		node2.Neighbors = append(node2.Neighbors, node1)
	}
}

func readLines(file string) (int, []string) {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lines := strings.Split(string(bytes), "\n")
	n, err := strconv.Atoi(lines[0])
	abortOnError(err)
	return n, lines[1:]
}
