package algo

import . "turing/ds"

func DFSRecursive(node NodeValue, graph Graph) {
	graph.LogVisit(node)
	graph.Visit(node)
	for _, neighbor := range graph.NeighborsOf(node) {
		if !graph.IsVisited(neighbor) {
			DFSRecursive(neighbor, graph)
		}
	}
}

func DFSIterative(node NodeValue, graph Graph) {
	nodeStack := make([]NodeValue, 0, graph.Size())
	nodeStack = append(nodeStack, graph.Get(node).Value())
	for len(nodeStack) != 0 {
		node := nodeStack[len(nodeStack)-1]
		nodeStack = nodeStack[:len(nodeStack)-1]
		graph.LogVisit(node)
		for _, neighbor := range graph.NeighborsOf(node) {
			if !graph.IsVisited(NodeValue(neighbor)) {
				graph.Visit(neighbor)
				nodeStack = append(nodeStack, NodeValue(neighbor))
			}
		}
	}
}
