package main

import (
	s "DIV-01/lem-in/solver"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	numberOfAnts := numberOfAnts()
	lines := readLines()
	// numberOfAnts := 1000
	// lines := []string{
	// 	"##start",
	// 	"1 23 3",
	// 	"2 16 7",
	// 	"#comment",
	// 	"3 16 3",
	// 	"4 16 5",
	// 	"5 9 3",
	// 	"6 1 5",
	// 	"7 4 8",
	// 	"##end",
	// 	"0 9 5",
	// 	"8 25 25",
	// 	"9 30 30",
	// 	"10 35 35",
	// 	"11 40 40",
	// 	"12 45 45",
	// 	"0-4",
	// 	"0-6",
	// 	"1-3",
	// 	"4-3",
	// 	"5-2",
	// 	"3-5",
	// 	"#another comment",
	// 	"4-2",
	// 	"2-1",
	// 	"7-6",
	// 	"7-2",
	// 	"7-4",
	// 	"6-5",
	// 	"1-8",
	// 	"8-9",
	// 	"9-10",
	// 	"10-11",
	// 	"11-12",
	// 	"12-0",
	// }
	mapOfNodes, startNode, endNode := parseLines(lines)
	var paths [][]string
	for range startNode.Neighbors {
		paths = append(paths, reverse(s.Solver(startNode, endNode, mapOfNodes))[1:])

	}

	lemin(numberOfAnts, paths, mapOfNodes)

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

func numberOfAnts() int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	number, err := strconv.Atoi(scanner.Text())
	abortOnError(err)
	return number
}
func readLines() []string {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		lines = append(lines, line)

	}
	err := scanner.Err()
	abortOnError(err)

	return lines
}
