package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	numberOfAnts := numberOfAnts()
	lines := readLines()
	parseLines(lines)

	fmt.Println(numberOfAnts)
}

type Node struct {
	Name      string
	X         int
	Y         int
	Visited   bool
	Neighbors []*Node
}

func abortOnError(err error) {
	if err != nil {
		fmt.Println("Normalno dannye vvodi, sumelek")
		os.Exit(1)
	}

}
func buildNodes(nodes [][]string) map[string]*Node {
	var n map[string]*Node = make(map[string]*Node)
	for _, v := range nodes {
		x, err := strconv.Atoi(v[1])
		abortOnError(err)
		y, err := strconv.Atoi(v[2])
		abortOnError(err)
		n[v[0]] = &Node{Name: v[0], X: x, Y: y}
	}
	return n
}
func parseLines(lines []string) {
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

	n := buildNodes(nodes)
	for i, v := range n {
		fmt.Println(i, *v)
	}
	fmt.Println(startNode)
	fmt.Println(endNode)

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
