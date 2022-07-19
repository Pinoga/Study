package ds

import (
	"fmt"
	"strings"
)

type GraphNode interface {
	Value() NodeValue
}

type Graph interface {
	NeighborsOf(NodeValue) []NodeValue
	LogVisit(NodeValue)
	Leave(NodeValue)
	Visit(NodeValue)
	IsVisited(NodeValue) bool
	Size() int
	Get(NodeValue) GraphNode
	New(int) Graph
	NewEdge(NodeValue, NodeValue)
	AddNodes(...GraphNode)
	Clear()
}

type GraphImpl struct {
	visited []bool
	adjList [][]NodeValue
	size    int
	nodes   map[NodeValue]GraphNode
	depth   int
	log     bool
}

func (nv NodeValue) Value() NodeValue {
	return nv
}

func (g *GraphImpl) logWithDepth(message string) {
	fmt.Print(strings.Repeat("\t", g.depth))
	fmt.Println(message)
}

func (g *GraphImpl) LogVisit(key NodeValue) {
	if g.log {
		g.logWithDepth(fmt.Sprintf("Visited %d", key))
	}
	g.depth++
}

func (g *GraphImpl) Visit(key NodeValue) {
	g.visited[key] = true
}

func (g *GraphImpl) Leave(key NodeValue) {
	g.depth--
	if g.log {
		g.logWithDepth(fmt.Sprintf("Finished %d", key))
	}
}

func (g *GraphImpl) IsVisited(key NodeValue) bool {
	return g.visited[key]
}

func (g *GraphImpl) Size() int {
	return g.size
}

func (g *GraphImpl) NewEdge(from NodeValue, to NodeValue) {
	g.adjList[from.Value()] = append(g.adjList[from.Value()], to.Value())

}

func (g *GraphImpl) Get(key NodeValue) GraphNode {
	return g.nodes[key]
}

func (g *GraphImpl) NeighborsOf(key NodeValue) []NodeValue {
	return g.adjList[key]
}

func (g *GraphImpl) New(size int) Graph {
	g.size = size
	g.nodes = make(map[NodeValue]GraphNode)

	for i := 0; i < size; i++ {
		g.nodes[NodeValue(i)] = NodeValue(i)
	}

	g.visited = make([]bool, size)
	g.adjList = make([][]NodeValue, size)
	g.log = true
	for nodeIndex := range g.adjList {
		g.adjList[nodeIndex] = make([]NodeValue, 0)
	}
	return g
}

func (g *GraphImpl) AddNodes(nodes ...GraphNode) {
	for _, node := range nodes {
		g.nodes[node.Value()] = node
	}
}

func (g *GraphImpl) Clear() {
	g.visited = make([]bool, g.size)
	g.depth = 0
}
