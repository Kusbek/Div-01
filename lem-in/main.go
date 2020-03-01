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
	// numberOfAnts := numberOfAnts()
	// lines := readLines()
	numberOfAnts := 25
	lines := []string{
		"##start",
		"1 23 3",
		"2 16 7",
		"#comment",
		"3 16 3",
		"4 16 5",
		"5 9 3",
		"6 1 5",
		"7 4 8",
		"##end",
		"0 9 5",
		"0-4",
		"0-6",
		"1-3",
		"4-3",
		"5-2",
		"3-5",
		"#another comment",
		"4-2",
		"2-1",
		"7-6",
		"7-2",
		"7-4",
		"6-5",
	}
	_, startNode, endNode := parseLines(lines)
	fmt.Println(numberOfAnts, startNode, endNode)
	s.Solver(numberOfAnts, startNode, endNode)

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
