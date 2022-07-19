package graph

import (
	"fmt"
	"strings"
)

var log bool = true
var depth int

var graph [][]int

var visited []bool

var components map[int]int

var traversalPath []int

func TraversalPath() []int {
	return traversalPath
}

func Components() map[int]int {
	return components
}

func NewGraph(adjList [][]int) {
	graph = adjList
	visited = make([]bool, len(adjList))
	components = make(map[int]int)
	traversalPath = make([]int, 0, len(adjList))
}

func visit(node int) {
	if log {
		logWithDepth(fmt.Sprintf("Visited %d", node))
	}
	depth++
	visited[node] = true
}
func endVisit(node int) {
	depth--
	if log {
		logWithDepth(fmt.Sprintf("Finished %d", node))
	}
}
func alreadyVisited(node int) {
	if log {
		logWithDepth(fmt.Sprintf("Already visited %d", node))
	}
}
func logWithDepth(message string) {
	fmt.Print(strings.Repeat("\t", depth))
	fmt.Println(message)
}
