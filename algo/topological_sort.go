package algo

import (
	. "turing/ds"
)

// func TopologicalSortIterative(graph Graph) []NodeValue {

// }
func TopologicalSortRecursive(graph Graph, node NodeValue, path []NodeValue, currSortIndex int) int {
	graph.Visit(node)
	graph.LogVisit(node)
	neighbors := graph.NeighborsOf(node)
	for _, neighbor := range neighbors {
		if !graph.IsVisited(neighbor) {
			currSortIndex = TopologicalSortRecursive(graph, neighbor, path, currSortIndex)
		}
	}

	path[currSortIndex] = node
	graph.Leave(node)
	return currSortIndex - 1
}
